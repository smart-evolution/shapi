package api

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"github.com/smart-evolution/shapi/services/agentsniffer"
	"github.com/smart-evolution/shapi/utils"
	"net/http"
)

const sniffAgentsHref string = "/api/sniffagents"

// CtrSniffAgents - api controller for sniffing agents
func CtrSniffAgents(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	data := struct {
		Devices string `json:"status"`
	}{
		"pending",
	}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": sniffAgentsHref,
		},
	}

	st := s.GetDataSource("state")

	state, ok := st.(state.IState)
	if !ok {
		utils.Log("Store should implement IState")
		handlers.HandleError(w, sniffAgentsHref, "controller store error", http.StatusInternalServerError)
		return
	}

	go agentsniffer.SniffAgents(state)
	embedded := map[string]string{}

	handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
}
