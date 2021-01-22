package agentconfigs

import (
	"encoding/json"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/domain/models/agent"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

// CtrAgentConfigPost -
func (c *Controller) CtrAgentConfigPost(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	agentID := opt.Params["agent"]
	href := "api/agentsConfig/" + agentID

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": href,
		},
	}

	var config agent.Config
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	decoder.Decode(&config)
	config.AgentID = agentID

	err := c.AgentUsecases.UpdateAgentConfigs(agentID, config)

	if err != nil {
		logger.Log("error while posting agentConfig")
		handlers.HandleError(w, href, "error while posting agentConfig", http.StatusInternalServerError)
		return
	}

	embedded := map[string]string{}
	handlers.HandleResponse(w, config, embedded, links, http.StatusOK)
}
