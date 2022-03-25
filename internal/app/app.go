package app

import (
	"context"

	"github.com/DanielTitkov/predictor/internal/configs"
	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/logger"
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

	// init app jobs, caches and preload data (if any)
	go app.UpdateSystemSummaryJob() // TODO: move to jobs?

	return &app, nil
}
