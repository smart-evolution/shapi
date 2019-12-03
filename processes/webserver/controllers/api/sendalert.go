package api

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"github.com/smart-evolution/shapi/utils"
	"net/http"
	"strconv"
)

const sendAlertHref string = "/api/sendalert"

// CtrSendAlert - api controller for sending alerts to agents
func CtrSendAlert(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	st := s.GetDataSource("state")

	state, ok := st.(state.IState)
	if !ok {
		utils.Log("store should implement IState")
		handlers.HandleError(w, sendAlertHref, "controller store error", http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		state.SetSendAlert(true)
	}

	data := struct {
		SendAlert string `json:"isAlerts"`
	}{
		strconv.FormatBool(state.SendAlert()),
	}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": sendAlertHref,
		},
	}

	embedded := map[string]string{}

	handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
}
