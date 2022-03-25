package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/DanielTitkov/predictor/internal/domain"

	"github.com/jfyne/live"
)

type (
	SystemSummaryInstance struct {
		Summary *domain.SystemSymmary
		Error   error
	}
)

func (h *Handler) NewSystemSummaryInstance(ctx context.Context, s live.Socket) *SystemSummaryInstance {
	m, ok := s.Assigns().(*SystemSummaryInstance)
	if !ok {
		summary, err := h.app.GetSystemSummary(ctx)
		return &SystemSummaryInstance{
			Summary: summary,
			Error:   err,
		}
	}

	return m
}

func (h *Handler) SystemSummary() live.Handler {
	t, err := template.ParseFiles(h.t+"layout.html", h.t+"system_summary.html")
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		return h.NewSystemSummaryInstance(ctx, s), nil
	})

	return lvh
}
