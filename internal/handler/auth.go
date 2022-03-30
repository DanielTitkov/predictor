package handler

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
)

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

	user, err := h.app.AuthenticateGothUser(req.Context(), &gu)
	if err != nil {
		h.log.Error("failed to create user", err)
		fmt.Fprintln(res, err)
		return
	}

	h.log.Debug("user authenticated", fmt.Sprintf("email: %s, provider: %s", user.Email, gu.Provider))

	// add or update session for user
	ses, err := h.app.CreateOrUpdateUserSession(req, user, true)
	if err != nil {
		h.log.Error("failed to create user session", err)
		fmt.Fprintln(res, err)
		return
	}

	h.log.Debug("user session refreshed", fmt.Sprintf("email: %s, sid: %s", user.Email, ses.SID))

	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}
