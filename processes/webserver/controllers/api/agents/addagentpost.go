package agents

import (
	"encoding/json"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"io/ioutil"
	"net/http"
)

type body struct {
	ID   string `json:"agentID"`
	Name string `json:"agentName"`
	IP   string `json:"agentIP"`
	Type string `json:"agentType"`
}

// CtrAdd - add new agent by IP
func (c *Controller) CtrAdd(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		logger.Log("error reading request body")
		http.Error(w, "error reading request body", http.StatusInternalServerError)
		return
	}

	var msg body
	err = json.Unmarshal(b, &msg)

	if err != nil {
		logger.Log("error parsing request body")
		http.Error(w, "error parsing request body", http.StatusInternalServerError)
		return
	}

	c.AgentUsecases.AddAgent(msg.ID, msg.Name, msg.IP, msg.Type)

	data := struct {
		Devices string `json:"message"`
	}{
		"agent added",
	}

	embedded := map[string]string{}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/agents/add",
		},
	}

	handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
}
