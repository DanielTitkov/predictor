package handler

import (
	"github.com/DanielTitkov/predictor/internal/app"
	"github.com/DanielTitkov/predictor/logger"
)

type (
	Handler struct {
		app *app.App
		log *logger.Logger
		t   string // template path
	}
)

func NewHandler(
	app *app.App,
	logger *logger.Logger,
	t string,
) *Handler {
	return &Handler{
		app: app,
		log: logger,
		t:   t,
	}
}
