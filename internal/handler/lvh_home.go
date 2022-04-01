package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/jfyne/live"
)

type (
	HomeInstance struct {
		*CommonInstance
		Summary                       *domain.SystemSymmary
		RandomFinishedChallenges      []*domain.Challenge
		RandomFinishedChallengesCount int
		RandomOngoingChallenges       []*domain.Challenge
		RandomOngoingChallengesCount  int
	}
)

func (h *Handler) NewHomeInstance(s live.Socket) *HomeInstance {
	m, ok := s.Assigns().(*HomeInstance)
	if !ok {
		return &HomeInstance{
			CommonInstance: h.NewCommon(s),
		}
	}

	return m
}

func (h *Handler) Home() live.Handler {
	t, err := template.ParseFiles(
		h.t+"layout.html",
		h.t+"home.html",
		h.t+"challenge_card.html",
		h.t+"system_summary.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewHomeInstance // NB: make sure constructor is correct
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
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		instance := h.NewHomeInstance(s)
		instance.fromContext(ctx)
		randomFinishedChallenges, err := h.app.GetRandomFinishedChallenges(ctx)
		if err != nil {
			instance.Error = err
		}
		instance.RandomFinishedChallenges = randomFinishedChallenges
		instance.RandomFinishedChallengesCount = len(randomFinishedChallenges)

		randomOngoingChallenges, err := h.app.GetRandomOngoingChallenges(ctx, instance.UserID)
		if err != nil {
			instance.Error = err
		}
		instance.RandomOngoingChallenges = randomOngoingChallenges
		instance.RandomOngoingChallengesCount = len(randomOngoingChallenges)

		summary, err := h.app.GetSystemSummary(ctx)
		if err != nil {
			instance.Error = err
		}
		instance.Summary = summary

		return instance, nil
	})

	return lvh
}
