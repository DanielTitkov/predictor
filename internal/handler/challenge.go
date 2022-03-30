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
		CommonInstance
		Challenge *domain.Challenge
	}
)

func (h *Handler) NewChallengeDetailsInstance(ctx context.Context, s live.Socket) *ChallengeDetailsInstance {
	m, ok := s.Assigns().(*ChallengeDetailsInstance)
	if !ok {
		return &ChallengeDetailsInstance{
			CommonInstance: CommonInstance{
				Env:     h.app.Cfg.Env,
				Session: fmt.Sprint(s.Session()),
				Error:   nil,
			},
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
		instance.User = UserFromCtx(ctx)
		challenge, err := h.app.GetChallengeByID(ctx, challengeID)
		if err != nil {
			instance.Error = err
		}
		instance.Challenge = challenge

		// ses, err := h.app.Store.Get(r, "email") // FIXME
		// fmt.Println("SES", ses.ID, ses.Name(), ses.Values, err)

		return instance, nil
	})

	return lvh
}
