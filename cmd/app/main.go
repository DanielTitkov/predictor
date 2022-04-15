package main

import (
	"context"
	"errors"
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

	var dbOptions []ent.Option

	if cfg.Env == "dev" {
		dbOptions = append(dbOptions, ent.Debug())
	}
	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI, dbOptions...)
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

	gothic.Store = store.Store
	goth.UseProviders(
		google.New(
			cfg.Auth.Google.Client,   // client
			cfg.Auth.Google.Secret,   // secret
			cfg.Auth.Google.Callback, // callback url
			"email", "profile",       // scopes
		),
	)

	h := handler.NewHandler(a, logger, "templates/")
	r := mux.NewRouter()
	r.Use(h.Middleware)
	r.NotFoundHandler = http.HandlerFunc(h.NotFoundRedirect)
	// main handler
	r.Handle("/challenge/{challengeID}", live.NewHttpHandler(store, h.ChallengeDetails()))
	r.Handle("/challenges", live.NewHttpHandler(store, h.ChallengeList()))
	r.Handle("/about", live.NewHttpHandler(store, h.About()))
	r.Handle("/profile", live.NewHttpHandler(store, h.Profile()))
	r.Handle("/404", live.NewHttpHandler(store, h.NotFound()))
	r.Handle("/", live.NewHttpHandler(store, h.Home()))

	// live scripts
	r.Handle("/live.js", live.Javascript{})
	r.Handle("/auto.js.map", live.JavascriptMap{})

	// auth
	r.HandleFunc("/auth/logout", h.Logout)
	r.HandleFunc("/auth/{provider}", h.BeginOAuth)
	r.HandleFunc("/auth/{provider}/callback", h.CompleteOAuth)

	// static
	r.HandleFunc("/favicon.ico", faviconHandler)
	r.HandleFunc("/static/css/styles.css", stylesHandler)

	// serve
	logger.Info("starting server", cfg.Server.GetAddress())
	log.Fatal(http.ListenAndServe(cfg.Server.GetAddress(), r))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/favicon.ico")
}

func stylesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/dist/css/styles.css")
}
