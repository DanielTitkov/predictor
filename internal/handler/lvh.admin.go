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
	eventAdminUpdatePage        = "challenge-list-update-page"
	eventAdminSelectPending     = "select-pending"
	eventAdminSelectUnpublished = "select-unpublished"
	eventAdminCreateNew         = "create-new"
	// params
	paramAdminPage = "page"
)

type (
	AdminInstance struct {
		*CommonInstance
		Challenges          []*domain.Challenge
		ChallengeCount      int
		FilterArgs          domain.FilterChallengesArgs
		Page                int
		MaxPage             int
		CreateChallengeForm bool
		TimeLayout          string
	}
)

func (ins *AdminInstance) withError(err error) *AdminInstance {
	ins.Error = err
	return ins
}

func (ins *AdminInstance) NextPage() int {
	if ins.Page >= ins.MaxPage {
		return ins.Page
	}
	return ins.Page + 1
}

func (ins *AdminInstance) PrevPage() int {
	if ins.Page == 1 {
		return ins.Page
	}
	return ins.Page - 1
}

func (ins *AdminInstance) updateChallenges(ctx context.Context, h *Handler) error {
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

func (h *Handler) NewAdminInstance(s live.Socket) *AdminInstance {
	m, ok := s.Assigns().(*AdminInstance)
	if !ok {
		return &AdminInstance{
			CommonInstance: h.NewCommon(s),
			Page:           1,
			FilterArgs: domain.FilterChallengesArgs{
				Pending:     true,
				Unpublished: false,
			},
			CreateChallengeForm: false,
			TimeLayout:          h.app.Cfg.App.DefaultTimeLayout,
		}
	}

	return m
}

func (h *Handler) Admin() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.admin.html",
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
		constructor := h.NewAdminInstance // NB: make sure constructor is correct
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
		instance := h.NewAdminInstance(s)
		instance.fromContext(ctx)

		if instance.User == nil || instance.UserID == uuid.Nil || !instance.User.Admin {
			s.Redirect(h.url404())
			return nil, nil
		}

		if err := instance.updateChallenges(ctx, h); err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	lvh.HandleEvent(eventAdminUpdatePage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		page := p.Int(paramAdminPage)
		instance := h.NewAdminInstance(s)
		instance.Page = page
		err := instance.updateChallenges(ctx, h)
		if err != nil {
			return instance.withError(err), nil
		}
		return instance, nil
	})

	lvh.HandleEvent(eventAdminSelectPending, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewAdminInstance(s)
		if instance.FilterArgs.Pending {
			return instance, nil
		}

		instance.Page = 1
		instance.FilterArgs.Pending = true
		instance.FilterArgs.Unpublished = false
		instance.CreateChallengeForm = false

		err := instance.updateChallenges(ctx, h)
		return instance.withError(err), nil
	})

	lvh.HandleEvent(eventAdminSelectUnpublished, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewAdminInstance(s)
		if instance.FilterArgs.Unpublished {
			return instance, nil
		}

		instance.Page = 1
		instance.FilterArgs.Pending = false
		instance.FilterArgs.Unpublished = true
		instance.CreateChallengeForm = false

		err := instance.updateChallenges(ctx, h)
		return instance.withError(err), nil
	})

	lvh.HandleEvent(eventAdminCreateNew, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewAdminInstance(s)

		instance.FilterArgs.Pending = false
		instance.FilterArgs.Unpublished = false
		instance.CreateChallengeForm = true

		return instance, nil
	})

	return lvh
}
