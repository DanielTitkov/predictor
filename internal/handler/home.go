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
		Session                  string
		RandomFinishedChallenges []*domain.Challenge
		RandomOngoingChallenges  []*domain.Challenge
		Error                    error
	}
)

func NewHomeInstance(s live.Socket) *HomeInstance {
	m, ok := s.Assigns().(*HomeInstance)
	if !ok {
		return &HomeInstance{
			Session: fmt.Sprint(s.Session()),
		}
	}

	return m
}

func (h *Handler) Home() live.Handler {
	t, err := template.ParseFiles(h.t+"layout.html", h.t+"home.html", h.t+"challenge_card.html")
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		instance := NewHomeInstance(s)
		randomFinishedChallenges, err := h.app.GetRandomFinishedChallenges(ctx)
		if err != nil {
			instance.Error = err
		}
		instance.RandomFinishedChallenges = randomFinishedChallenges

		randomOngoingChallenges, err := h.app.GetRandomOngoingChallenges(ctx)
		if err != nil {
			instance.Error = err
		}
		instance.RandomOngoingChallenges = randomOngoingChallenges

		return instance, nil
	})

	return lvh
}
