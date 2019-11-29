package api

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/utils"
	"net/http"
	"strconv"
)

// CtrAlerts - api controller for sending alerts to agents
func CtrAlerts(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	dfc := s.GetDataSource("state")

	st, ok := dfc.(state.IState)
	if !ok {
		utils.Log("store should implement IState")
		http.Error(w, "store should implement IState", http.StatusInternalServerError)
		return
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

	switch r.Method {
	case "OPTIONS":
		return
	case "GET":
		json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
		return
	case "POST":
		json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
		return
	}
}
