package agents

import (
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/domain/models/agent"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

// CtrAgentsPost -
func (c *Controller) CtrAgentsPost(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	agentID := opt.Params["agent"]

	href := "/api/agents/" + agentID

	ia, err := c.AgentUsecases.AgentByID(agentID)

	if err != nil {
		logger.Log("agent with id = " + agentID + " not found")
		handlers.HandleError(w, href, "agent not found", http.StatusNotFound)
		return
	}

	foundAgent, ok := ia.(*agent.Agent)

	if !ok {
		logger.Log("type assertion error")
		return
	}

	_, err = http.Get(foundAgent.IP)

	if err != nil {
		logger.Log("requesting agent with IP = " + foundAgent.IP + " failed")
		handlers.HandleError(w, href, "error contacting agent", http.StatusInternalServerError)
		return
	}
}
