package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/jfyne/live"
)

type (
	AboutInstance struct {
		CommonInstance
		Summary *domain.SystemSymmary
	}
)

func (h *Handler) NewAboutInstance(s live.Socket) *AboutInstance {
	m, ok := s.Assigns().(*AboutInstance)
	if !ok {
		return &AboutInstance{
			CommonInstance: CommonInstance{
				Env:     h.app.Cfg.Env,
				Session: fmt.Sprint(s.Session()),
				Error:   nil,
			},
		}
	}

	return m
}

func (h *Handler) About() live.Handler {
	t, err := template.ParseFiles(
		h.t+"layout.html",
		h.t+"about.html",
		h.t+"system_summary.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		instance := h.NewAboutInstance(s)
		instance.fromContext(ctx)

		summary, err := h.app.GetSystemSummary(ctx)
		if err != nil {
			instance.Error = err
		}
		instance.Summary = summary

		return instance, nil
	})

	return lvh
}
