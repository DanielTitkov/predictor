package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events

	// params
	paramsTaskDetailsTaskID = "taskID"
)

type (
	TaskDetailsInstance struct {
		CommonInstance
		Page    int
		MaxPage int
	}
)

func (t *TaskDetailsInstance) NextPage() int {
	if t.Page >= t.MaxPage {
		return t.Page
	}
	return t.Page + 1
}

func (t *TaskDetailsInstance) PrevPage() int {
	if t.Page == 1 {
		return t.Page
	}
	return t.Page - 1
}

// func (t *TaskDetailsInstance) updateTasks(ctx context.Context, h *Handler) {
// 	tasks, err := h.app.GetTasks(ctx, h.app.Cfg.App.DefaultTaskPageLimit, (t.Page-1)*h.app.Cfg.App.DefaultTaskPageLimit)
// 	if err != nil {
// 		t.Error = err
// 		return
// 	}

// 	t.Tasks = tasks
// }

func (h *Handler) NewTaskDetailsInstance(ctx context.Context, s live.Socket, taskID int) *TaskDetailsInstance {
	m, ok := s.Assigns().(*TaskDetailsInstance)
	if !ok {
		var nTasks int
		return &TaskDetailsInstance{
			CommonInstance: CommonInstance{
				Env:     h.app.Cfg.Env,
				Session: fmt.Sprint(s.Session()),
				Error:   nil,
			},
			Page:    1,
			MaxPage: int(math.Ceil(float64(nTasks) / float64(h.app.Cfg.App.DefaultChallengePageLimit))),
		}
	}

	return m
}

func (h *Handler) TaskDetails() live.Handler {
	t, err := template.ParseFiles(h.t+"layout.html", h.t+"task_details.html")
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		r := live.Request(ctx)
		taskIDStr, ok := mux.Vars(r)[paramsTaskDetailsTaskID]
		if !ok {
			return nil, errors.New("task id is required")
		}
		taskID, err := strconv.ParseInt(taskIDStr, 10, 64)
		if err != nil {
			return nil, err
		}
		instance := h.NewTaskDetailsInstance(ctx, s, int(taskID))
		instance.User, instance.UserID = UserFromCtx(ctx)

		// i.updateTasks(ctx, h)
		return instance, nil
	})

	// lvh.HandleEvent(eventTasksUpdatePage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	// 	page := p.Int(paramTasksPage)
	// 	v := url.Values{}
	// 	v.Add(paramTasksPage, fmt.Sprintf("%d", page))
	// 	s.PatchURL(v)
	// 	return s.Assigns(), nil
	// })

	return lvh
}
