package auth

import (
	"fmt"
	"net/http"

	"github.com/DanielTitkov/predictor/internal/app"
	"github.com/DanielTitkov/predictor/logger"
	"github.com/markbates/goth/gothic"
)

type Handler struct {
	app *app.App
	log *logger.Logger
}

func NewHandler(
	app *app.App,
	logger *logger.Logger,
) *Handler {
	return &Handler{
		app: app,
		log: logger,
	}
}

func (h *Handler) BeginOAuth(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func (h *Handler) CompleteOAuth(res http.ResponseWriter, req *http.Request) {
	gu, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		h.log.Error("failed to complete oauth", err)
		fmt.Fprintln(res, err)
		return
	}

	sesID, err := h.app.LiveSessionID(req) // FIXME
	fmt.Println("LIVE SESSION IN AUTH", sesID, err)

	user, err := h.app.AuthenticateGothUser(req.Context(), &gu)
	if err != nil {
		h.log.Error("failed to create user", err)
		fmt.Fprintln(res, err)
		return
	}

	h.log.Debug("user authenticated", fmt.Sprintf("email: %s, provider: %s", user.Email, gu.Provider))

	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}
