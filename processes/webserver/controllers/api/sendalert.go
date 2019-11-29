package api

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

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

	json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
}
