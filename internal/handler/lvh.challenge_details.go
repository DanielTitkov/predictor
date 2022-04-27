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
	eventAddPrediction                = "add-prediction"
	eventAddPredictionModal           = "add-prediction-modal"
	eventCloseModal                   = "close-modal"
	eventChallengeDetailsToggleProofs = "toggle-proofs"
	// params
	paramChallengeDetailsChallengeID = "challengeID"
	paramAddPredictionValue          = "addprediction"
	// params values
)

type (
	ChallengeDetailsInstance struct {
		*CommonInstance
		Challenge       *domain.Challenge
		ShowModal       bool
		ModalPrediction bool
		ShowProofs      bool
	}
)

func (h *Handler) NewChallengeDetailsInstance(s live.Socket) *ChallengeDetailsInstance {
	m, ok := s.Assigns().(*ChallengeDetailsInstance)
	if !ok {
		return &ChallengeDetailsInstance{
			CommonInstance: h.NewCommon(s, viewChallengeDetails),
			ShowModal:      false,
			ShowProofs:     false,
		}
	}

	return m
}

func (ins *ChallengeDetailsInstance) toggleProofs() {
	ins.ShowProofs = !ins.ShowProofs
}

func (h *Handler) ChallengeDetails() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.challenge_details.html",
		h.t+"part.challenge_details_scale.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewChallengeDetailsInstance // NB: make sure constructor is correct
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

		lvh.HandleEvent(eventCloseMessage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseMessage()
			return instance, nil
		})
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

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
		instance := h.NewChallengeDetailsInstance(s)
		instance.fromContext(ctx)
		challenge, err := h.app.GetChallengeByID(ctx, challengeID, instance.UserID)
		if err != nil {
			instance.Error = err
		}
		instance.Challenge = challenge

		return instance, nil
	})

	lvh.HandleEvent(eventAddPredictionModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeDetailsInstance(s)
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
		instance := h.NewChallengeDetailsInstance(s)
		instance.ShowModal = false
		return instance, nil
	})

	lvh.HandleEvent(eventChallengeDetailsToggleProofs, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeDetailsInstance(s)
		instance.toggleProofs()
		return instance, nil
	})

	lvh.HandleEvent(eventAddPrediction, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeDetailsInstance(s)
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
