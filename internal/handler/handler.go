package handler

import (
	"context"
	"fmt"

	"github.com/jfyne/live"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/app"
	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/DanielTitkov/predictor/logger"
)

const (
	// events (common)
	eventCloseAuthModals = "close-auth-modals"
	eventOpenLogoutModal = "open-logout-modal"
	eventOpenLoginModal  = "open-login-modal"
	eventCloseError      = "close-error-notification"
	// context
	userCtxKeyValue = "user"
)

type (
	Handler struct {
		app *app.App
		log *logger.Logger
		t   string // template path
	}

	CommonInstance struct {
		Env             string
		Session         string
		Error           error
		User            *domain.User
		UserID          uuid.UUID
		ShowLoginModal  bool
		ShowLogoutModal bool
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

func (h *Handler) NewCommon(s live.Socket) *CommonInstance {
	return &CommonInstance{
		Env:             h.app.Cfg.Env,
		Session:         fmt.Sprint(s.Session()),
		Error:           nil,
		ShowLoginModal:  false,
		ShowLogoutModal: false,
	}
}

func (c *CommonInstance) CloseAuthModals() {
	c.ShowLoginModal = false
	c.ShowLogoutModal = false
}

func (c *CommonInstance) OpenLoginModal() {
	c.ShowLoginModal = true
}

func (c *CommonInstance) OpenLogoutModal() {
	c.ShowLogoutModal = true
}

func (c *CommonInstance) CloseError() {
	c.Error = nil
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
