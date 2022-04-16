package handler

import (
	"context"
	"errors"
	"html/template"
	"log"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
// events
// params
// params values
)

type (
	ChallengeUpdateInstance struct {
		*CommonInstance
		Challenge       *domain.Challenge
		ShowModal       bool
		ModalPrediction bool
	}
)

func (ins *ChallengeUpdateInstance) withError(err error) *ChallengeUpdateInstance {
	ins.Error = err
	return ins
}

func (h *Handler) NewChallengeUpdateInstance(s live.Socket) *ChallengeUpdateInstance {
	m, ok := s.Assigns().(*ChallengeUpdateInstance)
	if !ok {
		return &ChallengeUpdateInstance{
			CommonInstance: h.NewCommon(s),
			ShowModal:      false,
		}
	}

	return m
}

func (h *Handler) ChallengeUpdate() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.challenge_update.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewChallengeUpdateInstance // NB: make sure constructor is correct
		// SAFE TO COPY
		lvh.HandleEvent(eventCloseAuthModals, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseAuthModals()
			return instance, nil
		})

		lvh.HandleEvent(eventOpenLogoutModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.OpenLogoutModal()
			return instance, nil
		})

		lvh.HandleEvent(eventOpenLoginModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.OpenLoginModal()
			return instance, nil
		})

		lvh.HandleEvent(eventCloseError, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseError()
			return instance, nil
		})
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		instance := h.NewChallengeUpdateInstance(s)
		instance.fromContext(ctx)

		if instance.User == nil || instance.UserID == uuid.Nil || !instance.User.Admin {
			s.Redirect(h.url404())
			return nil, nil
		}

		r := live.Request(ctx)
		challengeIDStr, ok := mux.Vars(r)[paramChallengeDetailsChallengeID]
		if !ok {
			return nil, errors.New("challenge id is required")
		}

		challengeID, err := uuid.Parse(challengeIDStr)
		if err != nil {
			return nil, err
		}

		challenge, err := h.app.GetChallengeByID(ctx, challengeID, instance.UserID)
		if err != nil {
			instance.Error = err
		}
		instance.Challenge = challenge

		return instance, nil
	})

	return lvh
}
