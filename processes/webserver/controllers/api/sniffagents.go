package api

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/smarthome/datasources/state"
	"github.com/smart-evolution/smarthome/services/agentsniffer"
	"github.com/smart-evolution/smarthome/utils"
	"net/http"
)

// CtrSniffAgents - api controller for sniffing agents
func CtrSniffAgents(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	data := struct {
		Devices string `json:"status"`
	}{
		"pending",
	}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/sniffagents",
		},
	}

	st := s.GetDataSource("state")

	state, ok := st.(state.IState)
	if !ok {
		utils.Log("Store should implement IState")
		return
	}

	go agentsniffer.SniffAgents(state)
	embedded := map[string]string{}

	json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
}
