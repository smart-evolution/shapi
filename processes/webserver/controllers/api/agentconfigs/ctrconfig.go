package agentconfigs

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/smarthome/datasources"
	"github.com/smart-evolution/smarthome/datasources/persistence"
	"github.com/smart-evolution/smarthome/models/agent"
	utl "github.com/smart-evolution/smarthome/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

// CtrAgentConfig - controller for agents list
func CtrAgentConfig(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	defer r.Body.Close()
	agentID := opt.Params["agent"]

	dfc := s.GetDataSource(datasources.Persistence)

	p, ok := dfc.(persistence.IPersistance)
	if !ok {
		utl.Log("Invalid store - should implement `IPersistance`")
		return
	}
	c := p.GetCollection("agentConfigs")

	var config agent.Config

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "api/agentsConfig/" + agentID,
		},
	}

	switch r.Method {
	case "GET":
		var list []agent.Config

		err := c.Find(bson.M{}).All(&list)

		if err != nil {
			utl.Log("AgentConfig not found err=", err)
			return
		}

		data := map[string]string{
			"count": strconv.Itoa(len(list)),
		}

		embedded := map[string]interface{}{
			"configs": list,
		}
		json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))

	case "POST":
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&config)
		config.AgentID = agentID
		_, err := c.Upsert(
			bson.M{"agentId": agentID},
			config,
		)
		if err != nil {
			utl.Log(err)
		}
		embedded := map[string]string{}
		json.NewEncoder(w).Encode(helpers.ServeHal(config, embedded, links))

	default:
	}
}
