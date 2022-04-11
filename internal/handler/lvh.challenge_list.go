package handler

import (
	"context"
	"html/template"
	"log"
	"math"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/jfyne/live"
)

const (
	// events
	eventChallengeListUpdatePage = "challenge-list-update-page"
	eventFilterChallenges        = "filter-challenges"
	// params
	paramChallengeListPage     = "page"
	paramChallengeListOngoing  = "ongoing"
	paramChallengeListFinished = "finished"
	paramChallengeListUnvoted  = "unvoted"
	// params value
)

type (
	ChallengeListInstance struct {
		*CommonInstance
		Challenges     []*domain.Challenge
		ChallengeCount int
		FilterArgs     domain.FilterChallengesArgs
		Page           int
		MaxPage        int
	}
)

func (ins *ChallengeListInstance) withError(err error) *ChallengeListInstance {
	ins.Error = err
	return ins
}

func (ins *ChallengeListInstance) NextPage() int {
	if ins.Page >= ins.MaxPage {
		return ins.Page
	}
	return ins.Page + 1
}

func (ins *ChallengeListInstance) PrevPage() int {
	if ins.Page == 1 {
		return ins.Page
	}
	return ins.Page - 1
}

func (ins *ChallengeListInstance) updateChallenges(ctx context.Context, h *Handler) error {
	ins.FilterArgs.Limit = h.app.Cfg.App.DefaultChallengePageLimit
	ins.FilterArgs.UserID = ins.UserID
	if ins.Page > 0 {
		ins.FilterArgs.Offset = (ins.Page - 1) * h.app.Cfg.App.DefaultChallengePageLimit
	} else {
		ins.FilterArgs.Offset = 0
	}

	chs, count, err := h.app.FilterChallenges(ctx, &ins.FilterArgs)
	if err != nil {
		return err
	}
	ins.Challenges = chs
	ins.ChallengeCount = count
	ins.MaxPage = int(math.Ceil(float64(count) / float64(h.app.Cfg.App.DefaultChallengePageLimit)))

	return nil
}

func (h *Handler) NewChallengeListInstance(s live.Socket) *ChallengeListInstance {
	m, ok := s.Assigns().(*ChallengeListInstance)
	if !ok {
		return &ChallengeListInstance{
			CommonInstance: h.NewCommon(s),
			Page:           1,
		}
	}

	return m
}

func (h *Handler) ChallengeList() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.challenge_list.html",
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
		constructor := h.NewChallengeListInstance // NB: make sure constructor is correct
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
		instance := h.NewChallengeListInstance(s)
		instance.fromContext(ctx)

		err := instance.updateChallenges(ctx, h)
		if err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	lvh.HandleEvent(eventChallengeListUpdatePage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		page := p.Int(paramChallengeListPage)
		instance := h.NewChallengeListInstance(s)
		instance.Page = page
		err := instance.updateChallenges(ctx, h)
		if err != nil {
			return instance.withError(err), nil
		}
		return instance, nil
	})

	lvh.HandleEvent(eventFilterChallenges, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewChallengeListInstance(s)

		instance.FilterArgs.Ongoing = p.Checkbox(paramChallengeListOngoing)
		instance.FilterArgs.Finished = p.Checkbox(paramChallengeListFinished)
		instance.FilterArgs.Unvoted = p.Checkbox(paramChallengeListUnvoted)

		err := instance.updateChallenges(ctx, h)
		return instance.withError(err), nil
	})

	return lvh
}
