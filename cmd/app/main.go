package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/DanielTitkov/predictor/cmd/app/prepare"
	"github.com/DanielTitkov/predictor/internal/app"
	"github.com/DanielTitkov/predictor/internal/configs"
	"github.com/DanielTitkov/predictor/internal/handler"
	"github.com/DanielTitkov/predictor/internal/repository/entgo"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent"
	"github.com/DanielTitkov/predictor/logger"
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
		if cfg.Debug.LogDBQueries {
			dbOptions = append(dbOptions, ent.Debug())
		}
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

	store := prepare.Store(cfg)

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
	r := prepare.Mux(cfg, store, h)

	httpServer := prepare.Server(cfg, r)
	httpServer.Addr = cfg.Server.GetAddress()
	logger.Info("starting http server", cfg.Server.GetAddress())
	log.Fatal(httpServer.ListenAndServe())
}
