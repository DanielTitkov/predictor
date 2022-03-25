package job

import (
	"github.com/DanielTitkov/predictor/internal/app"
	"github.com/DanielTitkov/predictor/internal/configs"
	"github.com/DanielTitkov/predictor/logger"
)

type Job struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func New(
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Job {
	return &Job{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
}
