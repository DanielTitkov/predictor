package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/url"

	"github.com/jfyne/live"
)

const (
	// events
	eventTasksUpdatePage = "tasks-update-page"
	// params
	paramTasksPage = "page"
)

type (
	TasksInstance struct {
		Session string
		Page    int
		Error   error
		MaxPage int
	}
)

func (t *TasksInstance) NextPage() int {
	if t.Page >= t.MaxPage {
		return t.Page
	}
	return t.Page + 1
}

func (t *TasksInstance) PrevPage() int {
	if t.Page == 1 {
		return t.Page
	}
	return t.Page - 1
}

func (t *TasksInstance) updateTasks(ctx context.Context, h *Handler) {
	// tasks, err := h.app.GetTasks(ctx, h.app.Cfg.App.DefaultTaskPageLimit, (t.Page-1)*h.app.Cfg.App.DefaultTaskPageLimit)
	// if err != nil {
	// 	t.Error = err
	// 	return
	// }

	// t.Tasks = tasks
}

func (h *Handler) NewTasksInstance(ctx context.Context, s live.Socket) *TasksInstance {
	m, ok := s.Assigns().(*TasksInstance)
	if !ok {
		var nTasks int
		return &TasksInstance{
			Session: fmt.Sprint(s.Session()),
			Page:    1,
			MaxPage: int(math.Ceil(float64(nTasks) / float64(h.app.Cfg.App.DefaultChallengePageLimit))),
		}
	}

	return m
}

func (h *Handler) Tasks() live.Handler {
	t, err := template.ParseFiles(h.t+"layout.html", h.t+"tasks.html")
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		i := h.NewTasksInstance(ctx, s)
		i.updateTasks(ctx, h)
		return i, nil
	})

	lvh.HandleParams(func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		page := p.Int(paramTasksPage)
		i := h.NewTasksInstance(ctx, s)
		i.Page = page
		i.updateTasks(ctx, h)
		return i, nil
	})

	lvh.HandleEvent(eventTasksUpdatePage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		page := p.Int(paramTasksPage)
		v := url.Values{}
		v.Add(paramTasksPage, fmt.Sprintf("%d", page))
		s.PatchURL(v)
		return s.Assigns(), nil
	})

	return lvh
}
