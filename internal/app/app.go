package app

import (
	"context"

	"github.com/DanielTitkov/predictor/internal/configs"
	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/logger"
	"github.com/google/uuid"
)

type (
	App struct {
		Cfg           configs.Config
		log           *logger.Logger
		repo          Repository
		systemSummary *domain.SystemSymmary
	}
	Repository interface {
		// challenge
		CreateOrUpdateChallengeByContent(context.Context, *domain.Challenge) (*domain.Challenge, error)
		GetChallengeByContent(context.Context, string) (*domain.Challenge, error)

		// user
		GetUserCount(context.Context) (int, error)
		GetUserByEmail(context.Context, string) (*domain.User, error)
		GetUserByID(context.Context, uuid.UUID) (*domain.User, error)
		CreateUser(context.Context, *domain.User) (*domain.User, error)

		// prediction
		CreatePrediction(context.Context, *domain.Prediction) (*domain.Prediction, error)
		GetPredictionByUserAndChallenge(ctx context.Context, userID, challengeID uuid.UUID) (*domain.Prediction, error)
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) (*App, error) {
	app := App{
		Cfg:  cfg,
		log:  logger,
		repo: repo,
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
