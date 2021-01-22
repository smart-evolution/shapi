package sniffagents

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

const sniffAgentsHref string = "/api/sniffagents"

// CtrSniffAgentsAll - api controller for sniffing agents
func (c *Controller) CtrSniffAgentsAll(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
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

	go c.UserUsecases.SniffAgents()
	embedded := map[string]string{}

	handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
}
