package api

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/smarthome/datasources/state"
	"github.com/smart-evolution/smarthome/utils"
	"net/http"
	"strconv"
)

// CtrSendAlert - api controller for sending alerts to agents
func CtrSendAlert(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	st := s.GetDataSource("state")

	state, ok := st.(state.IState)
	if !ok {
		utils.Log("Invalid store")
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
			"href": "/api/sendalert",
		},
	}

	embedded := map[string]string{}

	json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
}
