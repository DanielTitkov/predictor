package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/url"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	// eventEditChallengeValidate = "edit-validate"
	eventEditChallengeSubmit = "edit-submit"
	eventEditOutcomeValidate = "outcome-validate"
	eventEditOutcomeSubmit   = "outcome-submit"
	// params
	paramEditChallengeContent      = "content"
	paramEditChallengeDescription  = "description"
	paramEditChallengeStartTime    = "start-time"
	paramEditChallengeEndTime      = "end-time"
	paramEditChallengePublished    = "published"
	paramEditChallengeProofContent = "proof-content" // add -1 -2 etc.
	paramEditChallengeProofLink    = "proof-link"    // add -1 -2 etc.
	paramEditChallengeOutcome      = "outcome"

// params values
)

type (
	ChallengeUpdateInstance struct {
		*CommonInstance
		Challenge     *domain.Challenge
		ChallengeArgs domain.CreateChallengeArgs
		FormPrefill   domain.CreateChallengeArgs
		FormError     error
		OutcomeError  error
		TimeLayout    string
	}
)

func (ins *ChallengeUpdateInstance) withError(err error) *ChallengeUpdateInstance {
	ins.Error = err
	return ins
}

func (ins *ChallengeUpdateInstance) initArgs() {
	ins.ChallengeArgs = domain.CreateChallengeArgs{
		Type:        ins.Challenge.Type,
		Content:     ins.Challenge.Content,
		Description: ins.Challenge.Description,
		StartTime:   ins.Challenge.StartStr(),
		EndTime:     ins.Challenge.EndStr(),
		Published:   ins.Challenge.Published,
	}
}

func (h *Handler) NewChallengeUpdateInstance(s live.Socket) *ChallengeUpdateInstance {
	m, ok := s.Assigns().(*ChallengeUpdateInstance)
	if !ok {
		return &ChallengeUpdateInstance{
			CommonInstance: h.NewCommon(s),
			FormError:      nil,
			OutcomeError:   errors.New("provide outcome and proofs"),
			TimeLayout:     h.app.Cfg.App.DefaultTimeLayout,
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

		lvh.HandleEvent(eventCloseMessage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseMessage()
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

		instance.initArgs()

		return instance, nil
	})

	// lvh.HandleEvent(eventEditChallengeValidate, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	// 	fmt.Println("VALIDATE", p) // FIXME
	// 	instance := h.NewChallengeUpdateInstance(s)

	// 	instance.ChallengeArgs = editChallengeArgsFromParams(p, h.app.Cfg.App.DefaultTimeLayout)
	// 	instance.FormError = instance.ChallengeArgs.Validate()

	// 	return instance, nil
	// })

	lvh.HandleEvent(eventEditChallengeSubmit, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeUpdateInstance(s)

		instance.ChallengeArgs = editChallengeArgsFromParams(p, h.app.Cfg.App.DefaultTimeLayout)
		instance.FormError = instance.ChallengeArgs.Validate()
		if instance.FormError != nil {
			return instance, nil
		}

		_, err := h.app.UpdateChallengeByID(ctx, instance.Challenge.ID, &instance.ChallengeArgs)
		if err != nil {
			return instance.withError(err), nil
		}

		u, _ := url.Parse(fmt.Sprintf("/challenge/%s", instance.Challenge.ID))
		s.Redirect(u)
		return nil, nil
	})

	lvh.HandleEvent(eventEditOutcomeValidate, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeUpdateInstance(s)

		instance.OutcomeError = nil
		proofs := instance.proofsFromParams(p, h.app.Cfg.App.MinProofCount)
		if len(proofs) < h.app.Cfg.App.MinProofCount {
			instance.OutcomeError = fmt.Errorf("minimal proof count is %d, but got %d", h.app.Cfg.App.MinProofCount, len(proofs))
		}

		return instance, nil
	})

	lvh.HandleEvent(eventEditOutcomeSubmit, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeUpdateInstance(s)

		instance.OutcomeError = nil
		proofs := instance.proofsFromParams(p, h.app.Cfg.App.MinProofCount)
		if len(proofs) < h.app.Cfg.App.MinProofCount {
			instance.OutcomeError = fmt.Errorf("minimal proof count is %d, but got %d", h.app.Cfg.App.MinProofCount, len(proofs))
			return instance, nil
		}

		outcome := p.Checkbox(paramEditChallengeOutcome)
		err := h.app.SetChallengeOutcome(ctx, instance.Challenge.ID, outcome, proofs)
		if err != nil {
			return instance.withError(err), nil
		}

		u, _ := url.Parse(fmt.Sprintf("/challenge/%s", instance.Challenge.ID))
		s.Redirect(u)

		return instance, nil
	})

	return lvh
}

func editChallengeArgsFromParams(p live.Params, layout string) domain.CreateChallengeArgs {
	return domain.CreateChallengeArgs{
		Content:     p.String(paramEditChallengeContent),
		Description: p.String(paramEditChallengeDescription),
		StartTime:   p.String(paramEditChallengeStartTime),
		EndTime:     p.String(paramEditChallengeEndTime),
		Published:   p.Checkbox(paramEditChallengePublished),
		TimeLayout:  layout,
	}
}

func (ins *ChallengeUpdateInstance) proofsFromParams(p live.Params, proofCount int) []*domain.Proof {
	var proofs []*domain.Proof

	for i := 0; i < proofCount; i++ {
		proof := &domain.Proof{
			Content: p.String(fmt.Sprintf("%s-%d", paramEditChallengeProofContent, i+1)),
			Link:    p.String(fmt.Sprintf("%s-%d", paramEditChallengeProofLink, i+1)),
		}

		if proof.Content == "" || proof.Link == "" {
			continue
		}

		proofs = append(proofs, proof)
	}

	return proofs
}
