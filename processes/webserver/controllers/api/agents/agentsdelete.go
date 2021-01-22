package agents

import (
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

// CtrAgentsDelete - delete handler
func (c *Controller) CtrAgentsDelete(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	agentID := opt.Params["agent"]
	href := "/api/agents/" + agentID

	err := c.AgentUsecases.RemoveAgent(agentID)

	if err != nil {
		msg := "error deleting agent with ID = " + agentID
		logger.Log(msg)
		handlers.HandleError(w, href, msg, http.StatusInternalServerError)
	}
}
