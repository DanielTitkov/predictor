package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/jfyne/live"
)

const (
// events
// params
// paramChallengeListChallengeID = "challengeID"
// params values
)

type (
	ChallengeListInstance struct {
		*CommonInstance
		Challenges     []*domain.Challenge
		ChallengeCount int
	}
)

func (h *Handler) NewChallengeListInstance(s live.Socket) *ChallengeListInstance {
	m, ok := s.Assigns().(*ChallengeListInstance)
	if !ok {
		return &ChallengeListInstance{
			CommonInstance: h.NewCommon(s),
		}
	}

	return m
}

func (h *Handler) ChallengeList() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.challenge_list.html",
		h.t+"part.challenge_card.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewChallengeListInstance // NB: make sure constructor is correct
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
		instance := h.NewChallengeListInstance(s)
		instance.fromContext(ctx)

		return instance, nil
	})

	return lvh
}
