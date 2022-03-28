package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DanielTitkov/predictor/cmd/app/prepare"
	"github.com/DanielTitkov/predictor/internal/app"
	"github.com/DanielTitkov/predictor/internal/configs"
	"github.com/DanielTitkov/predictor/internal/handler"
	"github.com/DanielTitkov/predictor/internal/repository/entgo"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent"
	"github.com/DanielTitkov/predictor/logger"
	"github.com/gorilla/mux"
	"github.com/jfyne/live"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	log.Println("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	log.Println("loaded config")

	logger := logger.NewLogger(cfg.Env)
	defer logger.Sync()
	logger.Info("starting service", "")

	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	err = prepare.Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	repo := entgo.NewEntgoRepository(db, logger)

	store := live.NewCookieStore("go-live-session", []byte(cfg.Auth.Secret))
	store.Store.Options.SameSite = http.SameSiteLaxMode
	store.Store.MaxAge(cfg.Auth.Exp)
	store.Store.Options.Path = "/"
	store.Store.Options.HttpOnly = true
	store.Store.Options.Secure = !(cfg.Env == "dev")

	a, err := app.New(cfg, logger, repo, store.Store)
	if err != nil {
		logger.Fatal("failed to init app", err)
	}

	h := handler.NewHandler(
		a,
		logger,
		"templates/",
	)

	gothic.Store = store.Store

	// store1 := sessions.NewCookieStore([]byte(cfg.Auth.Secret))
	// store1.MaxAge(cfg.Auth.Exp)
	// store1.Options.Path = "/"
	// store1.Options.HttpOnly = true
	// store1.Options.Secure = !(cfg.Env == "dev")

	// gothic.Store = store1

	goth.UseProviders(
		google.New(
			cfg.Auth.Google.Client,   // client
			cfg.Auth.Google.Secret,   // secret
			cfg.Auth.Google.Callback, // callback url
			// scopes
			"email",
			"profile",
		),
	)

	r := mux.NewRouter()
	// Run the server.
	r.Handle("/", live.NewHttpHandler(store, h.Home()))
	r.Handle("/challenge/{challengeID}", live.NewHttpHandler(store, h.ChallengeDetails()))
	r.Handle("/tasks", live.NewHttpHandler(store, h.Tasks()))
	// live scripts
	r.Handle("/live.js", live.Javascript{})
	r.Handle("/auto.js.map", live.JavascriptMap{})
	// auth
	r.HandleFunc("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})
	r.HandleFunc("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		fmt.Printf("GOTH USER\n%+v", user) // FIXME
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
	})
	// favicon
	r.HandleFunc("/favicon.ico", faviconHandler)

	// serve
	logger.Info("starting server", cfg.Server.GetAddress())
	log.Fatal(http.ListenAndServe(cfg.Server.GetAddress(), r))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/favicon.ico")
}
