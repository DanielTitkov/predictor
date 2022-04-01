package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/jfyne/live"
)

func (a *App) CreateUserSession(req *http.Request, user *domain.User) (*domain.UserSession, error) {
	// get session sid for request
	sid, err := a.LiveSessionID(req)
	if err != nil {
		return nil, err
	}

	session := &domain.UserSession{
		SID:       sid,
		UserAgent: req.UserAgent(),
		IP:        req.RemoteAddr,
		UserID:    user.ID,
	}

	// create session record for user
	session, err = a.repo.CreateUserSession(req.Context(), session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (a *App) CreateOrUpdateUserSession(req *http.Request, user *domain.User, withOAuth bool, setActiveStatus bool) (*domain.UserSession, error) {
	if !withOAuth {
		return nil, errors.New("not allowed without oauth")
	}

	// get session sid for request
	sid, err := a.LiveSessionID(req)
	if err != nil {
		return nil, err
	}

	session := &domain.UserSession{
		SID:       sid,
		UserAgent: req.UserAgent(),
		IP:        req.RemoteAddr,
		UserID:    user.ID,
		Active:    setActiveStatus,
	}

	// create session record for user
	session, err = a.repo.CreateOrUpdateUserSession(req.Context(), session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (a *App) GetUserBySession(req *http.Request) (*domain.User, error) {
	// get session sid for request
	sid, err := a.LiveSessionID(req)
	if err != nil {
		return nil, err
	}

	session := &domain.UserSession{
		SID:       sid,
		UserAgent: req.UserAgent(),
		IP:        req.RemoteAddr,
	}

	// check if session saved for some user
	registered, err := a.repo.IfSessionRegistered(req.Context(), session)
	if err != nil {
		return nil, err
	}

	// proceed without user
	if !registered {
		return nil, nil
	}

	// retrieve user and add to context
	user, err := a.repo.GetUserBySession(req.Context(), session)
	if err != nil {
		return nil, err
	}

	// update session activity
	err = a.repo.UpdateUserSessionLastActivityBySID(req.Context(), sid)
	if err != nil {
		return nil, err
	}

	return user, nil
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
