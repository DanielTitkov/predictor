package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"strconv"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	eventAddPrediction      = "add-prediction"
	eventAddPredictionModal = "add-prediction-modal"
	eventCloseModal         = "close-modal"
	// params
	paramChallengeDetailsChallengeID = "challengeID"
	paramAddPredictionValue          = "addprediction"
	// params values
)

type (
	ChallengeDetailsInstance struct {
		CommonInstance
		Challenge       *domain.Challenge
		ShowModal       bool
		ModalPrediction bool
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
			ShowModal: false,
		}
	}

	return m
}

func (h *Handler) ChallengeDetails() live.Handler {
	t, err := template.ParseFiles(
		h.t+"layout.html",
		h.t+"challenge_details.html",
		h.t+"challenge_details_scale.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		r := live.Request(ctx)
		challengeIDStr, ok := mux.Vars(r)[paramChallengeDetailsChallengeID]
		if !ok {
			return nil, errors.New("challenge id is required")
		}

		challengeID, err := uuid.Parse(challengeIDStr)
		if err != nil {
			return nil, err
		}
		instance := h.NewChallengeDetailsInstance(ctx, s)
		instance.User, instance.UserID = UserFromCtx(ctx)
		challenge, err := h.app.GetChallengeByID(ctx, challengeID, instance.UserID)
		if err != nil {
			instance.Error = err
		}
		instance.Challenge = challenge

		return instance, nil
	})

	lvh.HandleEvent(eventAddPredictionModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeDetailsInstance(ctx, s)
		predictionValue, err := strconv.ParseBool(p.String(paramAddPredictionValue))
		if err != nil {
			instance.Error = err
			return instance, nil
		}
		instance.ModalPrediction = predictionValue
		instance.ShowModal = true
		return instance, nil
	})

	lvh.HandleEvent(eventCloseModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeDetailsInstance(ctx, s)
		instance.ShowModal = false
		return instance, nil
	})

	lvh.HandleEvent(eventAddPrediction, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeDetailsInstance(ctx, s)
		predictionValue, err := strconv.ParseBool(p.String(paramAddPredictionValue))
		if err != nil {
			instance.Error = err
			return instance, nil
		}

		if instance.User == nil {
			instance.Error = fmt.Errorf("you must be a registered user to add prediction")
			return instance, nil
		}

		// TODO: add session to meta
		pred := &domain.Prediction{
			Prognosis:   predictionValue,
			UserID:      instance.User.ID,
			ChallengeID: instance.Challenge.ID,
		}

		pred, err = h.app.CreatePrediction(ctx, pred)
		if err != nil {
			instance.Error = err
			return instance, nil
		}

		instance.Challenge.UserPrediction = pred
		instance.Challenge.Predictions = append(instance.Challenge.Predictions, pred)

		return instance, nil
	})

	return lvh
}
