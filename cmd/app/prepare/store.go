package prepare

import (
	"net/http"

	"github.com/DanielTitkov/predictor/internal/configs"
	"github.com/jfyne/live"
)

func Store(cfg configs.Config) *live.CookieStore {
	store := live.NewCookieStore("go-live-session", []byte(cfg.Auth.Secret))
	store.Store.Options.SameSite = http.SameSiteLaxMode
	store.Store.MaxAge(cfg.Auth.Exp)
	store.Store.Options.Path = "/"
	store.Store.Options.HttpOnly = true
	store.Store.Options.Secure = !(cfg.Env == "dev")
	return store
}
