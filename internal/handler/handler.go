package handler

import (
	"context"

	"github.com/DanielTitkov/predictor/internal/app"
	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/logger"
)

const (
	userCtxKeyValue = "user"
)

type (
	Handler struct {
		app *app.App
		log *logger.Logger
		t   string // template path
	}

	CommonInstance struct {
		Env     string
		Session string
		Error   error
		User    *domain.User
	}

	contextKey struct {
		name string
	}
)

var userCtxKey = &contextKey{userCtxKeyValue}

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

func UserFromCtx(ctx context.Context) *domain.User {
	user, ok := ctx.Value(userCtxKey).(*domain.User)
	if !ok {
		return nil
	}
	return user
}
