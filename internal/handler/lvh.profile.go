package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/DanielTitkov/predictor/internal/domain"
	"github.com/google/uuid"

	"github.com/jfyne/live"
)

type (
	ProfileInstance struct {
		*CommonInstance
		Summary     *domain.SystemSymmary
		UserSummary *domain.UserSummary
	}
)

func (ins *ProfileInstance) withError(err error) *ProfileInstance {
	ins.Error = err
	return ins
}

func (h *Handler) NewProfileInstance(s live.Socket) *ProfileInstance {
	m, ok := s.Assigns().(*ProfileInstance)
	if !ok {
		return &ProfileInstance{
			CommonInstance: h.NewCommon(s),
		}
	}

	return m
}

func (h *Handler) Profile() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.profile.html",
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

		return instance, nil
	})

	return lvh
}
