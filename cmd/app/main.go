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

	a, err := app.New(cfg, logger, repo)
	if err != nil {
		logger.Fatal("failed to init app", err)
	}

	h := handler.NewHandler(
		a,
		logger,
		"templates/",
	)

	r := mux.NewRouter()
	// Run the server.
	r.Handle("/", live.NewHttpHandler(live.NewCookieStore("session-name", []byte(cfg.Auth.Secret)), h.Home()))
	r.Handle("/summary", live.NewHttpHandler(live.NewCookieStore("session-name", []byte(cfg.Auth.Secret)), h.SystemSummary()))
	r.Handle("/challenge/{challengeID}", live.NewHttpHandler(live.NewCookieStore("session-name", []byte(cfg.Auth.Secret)), h.ChallengeDetails()))
	r.Handle("/tasks", live.NewHttpHandler(live.NewCookieStore("session-name", []byte(cfg.Auth.Secret)), h.Tasks()))
	// live scripts
	r.Handle("/live.js", live.Javascript{})
	r.Handle("/auto.js.map", live.JavascriptMap{})
	// favicon
	r.HandleFunc("/favicon.ico", faviconHandler)
	// serve
	log.Fatal(http.ListenAndServe(cfg.Server.GetAddress(), r))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/favicon.ico")
}
