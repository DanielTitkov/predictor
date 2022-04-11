package handler

import (
	"context"
	"html/template"
	"log"
	"math"

	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/google/uuid"

	"github.com/jfyne/live"
)

const (
	// events
	eventUserProfileUpdatePage = "challenge-list-update-page"
	eventProfileSelectOngoing  = "select-ongoing"
	eventProfileSelectFinished = "select-finished"
	// params
	paramUserProfilePage     = "page"
	paramUserProfileOngoing  = "ongoing"
	paramUserProfileFinished = "finished"
	// params value
)

type (
	ProfileInstance struct {
		*CommonInstance
		Summary        *domain.SystemSymmary
		UserSummary    *domain.UserSummary
		Challenges     []*domain.Challenge
		ChallengeCount int
		FilterArgs     domain.FilterChallengesArgs
		Page           int
		MaxPage        int
	}
)

func (ins *ProfileInstance) withError(err error) *ProfileInstance {
	ins.Error = err
	return ins
}

func (ins *ProfileInstance) NextPage() int {
	if ins.Page >= ins.MaxPage {
		return ins.Page
	}
	return ins.Page + 1
}

func (ins *ProfileInstance) PrevPage() int {
	if ins.Page == 1 {
		return ins.Page
	}
	return ins.Page - 1
}

func (ins *ProfileInstance) updateChallenges(ctx context.Context, h *Handler) error {
	ins.FilterArgs.Limit = h.app.Cfg.App.DefaultChallengePageLimit
	ins.FilterArgs.UserID = ins.UserID
	if ins.Page > 0 {
		ins.FilterArgs.Offset = (ins.Page - 1) * h.app.Cfg.App.DefaultChallengePageLimit
	} else {
		ins.FilterArgs.Offset = 0
	}

	chs, count, err := h.app.FilterUserChallenges(ctx, &ins.FilterArgs)
	if err != nil {
		return err
	}
	ins.Challenges = chs
	ins.ChallengeCount = count
	ins.MaxPage = int(math.Ceil(float64(count) / float64(h.app.Cfg.App.DefaultChallengePageLimit)))

	return nil
}

func (h *Handler) NewProfileInstance(s live.Socket) *ProfileInstance {
	m, ok := s.Assigns().(*ProfileInstance)
	if !ok {
		return &ProfileInstance{
			CommonInstance: h.NewCommon(s),
			Page:           1,
			FilterArgs: domain.FilterChallengesArgs{
				Ongoing:  false,
				Finished: true,
			},
		}
	}

	return m
}

func (h *Handler) Profile() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.profile.html",
		h.t+"part.challenge_list_item.html",
		h.t+"part.challenge_list_pagination.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewProfileInstance // NB: make sure constructor is correct
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
		instance := h.NewProfileInstance(s)
		instance.fromContext(ctx)

		if instance.User == nil || instance.UserID == uuid.Nil {
			s.Redirect(h.url404())
			return nil, nil
		}

		userSummary, err := h.app.GetUserSummary(ctx, instance.UserID)
		if err != nil {
			return instance.withError(err), nil
		}
		instance.UserSummary = userSummary

		if err := instance.updateChallenges(ctx, h); err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	lvh.HandleEvent(eventUserProfileUpdatePage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		page := p.Int(paramUserProfilePage)
		instance := h.NewProfileInstance(s)
		instance.Page = page
		err := instance.updateChallenges(ctx, h)
		if err != nil {
			return instance.withError(err), nil
		}
		return instance, nil
	})

	lvh.HandleEvent(eventProfileSelectOngoing, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewProfileInstance(s)
		if instance.FilterArgs.Ongoing {
			return instance, nil
		}

		instance.Page = 1
		instance.FilterArgs.Ongoing = true
		instance.FilterArgs.Finished = false

		err := instance.updateChallenges(ctx, h)
		return instance.withError(err), nil
	})

	lvh.HandleEvent(eventProfileSelectFinished, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewProfileInstance(s)
		if instance.FilterArgs.Finished {
			return instance, nil
		}

		instance.Page = 1
		instance.FilterArgs.Finished = true
		instance.FilterArgs.Ongoing = false

		err := instance.updateChallenges(ctx, h)
		return instance.withError(err), nil
	})

	return lvh
}
