package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/jfyne/live"
)

type (
	HomeInstance struct {
		CommonInstance
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
			CommonInstance: CommonInstance{
				Env:     h.app.Cfg.Env,
				Session: fmt.Sprint(s.Session()),
				Error:   nil,
			},
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
