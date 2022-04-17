package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	eventEditChallengeValidate = "edit-validate"
	eventEditChallengeSubmit   = "edit-submit"
	// params
	paramEditChallengeContent     = "content"
	paramEditChallengeDescription = "description"
	paramEditChallengeStartTime   = "start-time"
	paramEditChallengeEndTime     = "end-time"
	paramEditChallengePublished   = "published"

// params values
)

type (
	ChallengeUpdateInstance struct {
		*CommonInstance
		Challenge     *domain.Challenge
		ChallengeArgs domain.CreateChallengeArgs
		FormError     error
		TimeLayout    string
	}
)

func (ins *ChallengeUpdateInstance) withError(err error) *ChallengeUpdateInstance {
	ins.Error = err
	return ins
}

func (ins *ChallengeUpdateInstance) updateChallengeFromArgs() error {
	ins.Challenge.Content = ins.ChallengeArgs.Content
	ins.Challenge.Description = ins.ChallengeArgs.Description
	ins.Challenge.Published = ins.ChallengeArgs.Published

	startTime, err := time.Parse(ins.ChallengeArgs.TimeLayout, ins.ChallengeArgs.StartTime)
	if err != nil {
		return fmt.Errorf("failed to parse start time: %s", err)
	}

	endTime, err := time.Parse(ins.ChallengeArgs.TimeLayout, ins.ChallengeArgs.EndTime)
	if err != nil {
		return fmt.Errorf("failed to parse end time: %s", err)
	}

	ins.Challenge.StartTime = startTime
	ins.Challenge.EndTime = endTime

	return nil
}

func (h *Handler) NewChallengeUpdateInstance(s live.Socket) *ChallengeUpdateInstance {
	m, ok := s.Assigns().(*ChallengeUpdateInstance)
	if !ok {
		return &ChallengeUpdateInstance{
			CommonInstance: h.NewCommon(s),
			FormError:      nil,
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

	lvh.HandleEvent(eventEditChallengeValidate, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeUpdateInstance(s)

		instance.ChallengeArgs = editChallengeArgsFromParams(p, h.app.Cfg.App.DefaultTimeLayout, instance.UserID)
		instance.FormError = instance.ChallengeArgs.Validate()
		err := instance.updateChallengeFromArgs()
		if err != nil {
			instance.FormError = err
		}

		return instance, nil
	})

	return lvh
}

func editChallengeArgsFromParams(p live.Params, layout string, userID uuid.UUID) domain.CreateChallengeArgs {
	return domain.CreateChallengeArgs{
		Type:        domain.ChallengeTypeBool,
		Outcome:     nil,
		Content:     p.String(paramEditChallengeContent),
		Description: p.String(paramEditChallengeDescription),
		StartTime:   p.String(paramEditChallengeStartTime),
		EndTime:     p.String(paramEditChallengeEndTime),
		Published:   p.Checkbox(paramEditChallengePublished),
		TimeLayout:  layout,
		AuthorID:    userID,
	}
}
