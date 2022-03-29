package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DanielTitkov/predictor/internal/configs"
	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/logger"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jfyne/live"
)

type (
	App struct {
		Cfg           configs.Config
		log           *logger.Logger
		repo          Repository
		systemSummary *domain.SystemSymmary
		Store         sessions.Store
	}
	Repository interface {
		// challenge
		CreateOrUpdateChallengeByContent(context.Context, *domain.Challenge) (*domain.Challenge, error)
		GetChallengeByContent(context.Context, string) (*domain.Challenge, error)
		GetChallengeByID(context.Context, uuid.UUID) (*domain.Challenge, error)
		GetRandomFinishedChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error)
		GetRandomOngoingChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error)

		// user
		IfEmailRegistered(context.Context, string) (bool, error)
		GetUserByEmail(context.Context, string) (*domain.User, error)
		GetUserByID(context.Context, uuid.UUID) (*domain.User, error)
		CreateUser(context.Context, *domain.User) (*domain.User, error)

		// prediction
		CreatePrediction(context.Context, *domain.Prediction) (*domain.Prediction, error)
		GetPredictionByUserAndChallenge(ctx context.Context, userID, challengeID uuid.UUID) (*domain.Prediction, error)

		// for system summary
		GetChallengeCount(ctx context.Context) (int, error)
		GetOngoingChallengeCount(ctx context.Context) (int, error)
		GetFinishedChallengeCount(ctx context.Context) (int, error)
		GetUserCount(ctx context.Context) (int, error)
		GetPredictionCount(ctx context.Context) (int, error)
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
	store sessions.Store,
) (*App, error) {
	app := App{
		Cfg:   cfg,
		log:   logger,
		repo:  repo,
		Store: store,
	}

	err := app.loadChallengePresets()
	if err != nil {
		return nil, err
	}

	err = app.loadUserPresets()
	if err != nil {
		return nil, err
	}

	err = app.loadPredictionPresets()
	if err != nil {
		return nil, err
	}

	app.log.Info("finished loading presets", "")

	// init app jobs, caches and preload data (if any)
	go app.UpdateSystemSummaryJob() // TODO: move to jobs?

	return &app, nil
}

func (a *App) LiveSessionID(req *http.Request) (string, error) {
	ses, err := a.Store.Get(req, "go-live-session")
	if err != nil {
		return "", err
	}

	lsI := ses.Values["_ls"]
	ls, ok := lsI.(live.Session)
	if !ok {
		return "", fmt.Errorf("expected to get live.Session but got %T", lsI)
	}

	idI := ls["_lsid"]
	id, ok := idI.(string)
	if !ok {
		return "", fmt.Errorf("expected to get string but got %T", idI)
	}
	return id, nil
}
