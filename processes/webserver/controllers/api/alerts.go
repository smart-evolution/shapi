package api

import (
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
	"strconv"
)

// CtrAlerts - api controller for sending alerts to agents
func CtrAlerts(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	dfc := s.GetDataSource("state")

	st, ok := dfc.(state.IState)
	if !ok {
		logger.Log("store should implement IState")
		http.Error(w, "store should implement IState", http.StatusInternalServerError)
		return
	}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/alerts",
		},
	}

	embedded := map[string]string{}

	switch r.Method {
	case "OPTIONS":
		return
	case "GET":
		data := struct {
			IsAlerts string `json:"isAlerts"`
		}{
			strconv.FormatBool(st.IsAlerts()),
		}
		handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
		return
	case "POST":
		st.SetIsAlerts(!st.IsAlerts())

		data := struct {
			IsAlerts string `json:"isAlerts"`
		}{
			strconv.FormatBool(st.IsAlerts()),
		}
		handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
		return
	}
}
