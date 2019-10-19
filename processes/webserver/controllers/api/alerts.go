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

// CtrAlerts - api controller for sending alerts to agents
func CtrAlerts(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dfc := s.GetDataSource("state")

	st, ok := dfc.(state.IState)
	if !ok {
		utils.Log("Invalid store")
		return
	}

	if r.Method == "POST" {
		st.SetIsAlerts(!st.IsAlerts())
	}

	data := struct {
		IsAlerts string `json:"isAlerts"`
	}{
		strconv.FormatBool(st.IsAlerts()),
	}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/alerts",
		},
	}

	embedded := map[string]string{}

	w.Header().Set("Access-Control-Allow-Origin", "http://shpanel.xyz")
	json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
}
