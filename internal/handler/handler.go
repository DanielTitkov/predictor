package handler

import (
	"context"

	"github.com/google/uuid"

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
		UserID  uuid.UUID
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

func UserFromCtx(ctx context.Context) (*domain.User, uuid.UUID) {
	user, ok := ctx.Value(userCtxKey).(*domain.User)
	if !ok {
		return nil, uuid.Nil
	}
	if user == nil {
		return nil, uuid.Nil
	}
	return user, user.ID
}

func (c *CommonInstance) fromContext(ctx context.Context) {
	user, userID := UserFromCtx(ctx)
	c.User = user
	c.UserID = userID
}
