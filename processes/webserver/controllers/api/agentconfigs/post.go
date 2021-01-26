package agentconfigs

import (
	"encoding/json"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/domain/models/agent"
	"net/http"
)

// CtrAgentConfigPost - post handler
func (c *Controller) CtrAgentConfigPost(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	defer r.Body.Close()

	agentID := opt.Params["agent"]
	href := "api/agentsConfig/" + agentID

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": href,
		},
	}

	var config agent.Config
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&config)

	if err != nil {
		logger.Log("error while decoding agentConfig")
		c.HandleError(w, href, "error while decoding agentConfig", http.StatusInternalServerError)
		return
	}

	config.AgentID = agentID
	err = c.AgentUsecases.UpdateAgentConfigs(agentID, config)

	if err != nil {
		logger.Log("error while posting agentConfig")
		c.HandleError(w, href, "error while posting agentConfig", http.StatusInternalServerError)
		return
	}

	embedded := map[string]string{}
	c.HandleJSONResponse(w, config, embedded, links, http.StatusOK)
}
