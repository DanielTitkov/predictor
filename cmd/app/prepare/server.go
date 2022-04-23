package prepare

import (
	"net/http"

	"github.com/DanielTitkov/predictor/internal/configs"
	"github.com/DanielTitkov/predictor/internal/handler"
	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

func Server(cfg configs.Config, store live.HttpSessionStore, h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.Use(h.Middleware)
	r.NotFoundHandler = http.HandlerFunc(h.NotFoundRedirect)
	// main handler
	r.Handle("/challenge/{challengeID}/edit", live.NewHttpHandler(store, h.ChallengeUpdate()))
	r.Handle("/challenge/{challengeID}", live.NewHttpHandler(store, h.ChallengeDetails()))
	r.Handle("/challenges", live.NewHttpHandler(store, h.ChallengeList()))
	r.Handle("/about", live.NewHttpHandler(store, h.About()))
	r.Handle("/profile", live.NewHttpHandler(store, h.Profile()))
	r.Handle("/admin", live.NewHttpHandler(store, h.Admin()))
	r.Handle("/404", live.NewHttpHandler(store, h.NotFound()))
	r.Handle("/", live.NewHttpHandler(store, h.Home()))

	// live scripts
	r.Handle("/live.js", live.Javascript{})
	r.Handle("/auto.js.map", live.JavascriptMap{})

	// auth
	r.HandleFunc("/auth/logout", h.Logout)
	r.HandleFunc("/auth/{provider}", h.BeginOAuth)
	r.HandleFunc("/auth/{provider}/callback", h.CompleteOAuth)

	// static
	r.HandleFunc("/favicon.ico", faviconHandler)
	r.HandleFunc("/static/css/styles.css", stylesHandler)

	return r
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/favicon.ico")
}

func stylesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/dist/css/styles.css")
}
