package main

import (
	"context"
	"crypto/tls"
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
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/crypto/acme/autocert"

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

	var httpsServer *http.Server
	var certManager *autocert.Manager

	if cfg.Env != "dev" {
		httpsServer = prepare.Server(cfg, r)
		certManager := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("predictor.live"),
			Cache:      autocert.DirCache("certs"),
		}
		httpsServer.Addr = cfg.Server.GetAddress(true)
		httpsServer.TLSConfig = &tls.Config{GetCertificate: certManager.GetCertificate}

		go func() {
			logger.Info("starting https server on %s", cfg.Server.GetAddress(true))
			err := httpsServer.ListenAndServeTLS("", "")
			if err != nil {
				log.Fatalf("httpsServer.ListendAndServeTLS() failed with %s", err)
			}
		}()
	}

	httpServer := prepare.Server(cfg, r)
	if certManager != nil {
		httpServer.Handler = certManager.HTTPHandler(httpServer.Handler)
	}
	httpServer.Addr = cfg.Server.GetAddress(false)
	logger.Info("starting http server", cfg.Server.GetAddress(false))
	log.Fatal(httpServer.ListenAndServe())
}
