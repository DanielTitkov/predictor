package handler

import (
	"context"
	"errors"
	"fmt"
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
	paramsChallengeDetailsChallengeID = "challengeID"
)

type (
	ChallengeDetailsInstance struct {
		Session   string
		Challenge *domain.Challenge
		Error     error
	}
)

func (h *Handler) NewChallengeDetailsInstance(ctx context.Context, s live.Socket) *ChallengeDetailsInstance {
	m, ok := s.Assigns().(*ChallengeDetailsInstance)
	if !ok {
		return &ChallengeDetailsInstance{
			Session: fmt.Sprint(s.Session()),
			Error:   nil,
		}
	}

	return m
}

func (h *Handler) ChallengeDetails() live.Handler {
	t, err := template.ParseFiles(h.t+"layout.html", h.t+"challenge_details.html")
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		r := live.Request(ctx)
		challengeIDStr, ok := mux.Vars(r)[paramsChallengeDetailsChallengeID]
		if !ok {
			return nil, errors.New("challenge id is required")
		}

		challengeID, err := uuid.Parse(challengeIDStr)
		if err != nil {
			return nil, err
		}
		instance := h.NewChallengeDetailsInstance(ctx, s)
		challenge, err := h.app.GetChallengeByID(ctx, challengeID)
		if err != nil {
			instance.Error = err
		}
		instance.Challenge = challenge
		return instance, nil
	})

	return lvh
}
