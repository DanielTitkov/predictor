package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/jfyne/live"
)

type (
	HomeInstance struct {
		Session string
	}
)

func NewHomeInstance(s live.Socket) *HomeInstance {
	m, ok := s.Assigns().(*HomeInstance)
	if !ok {
		return &HomeInstance{
			Session: fmt.Sprint(s.Session()),
		}
	}

	return m
}

func (h *Handler) Home() live.Handler {
	t, err := template.ParseFiles(h.t+"layout.html", h.t+"home.html")
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		return NewHomeInstance(s), nil
	})

	return lvh
}
