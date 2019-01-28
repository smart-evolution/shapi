package api

import (
    "net/http"
    "github.com/smart-evolution/smarthome/datasources/persistence"
    "github.com/smart-evolution/smarthome/models/agent"
    utl "github.com/smart-evolution/smarthome/utils"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
    "github.com/coda-it/gowebserver/helpers"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/coda-it/gowebserver/store"
)

// CtrAgentEdit - controller for agents list
func CtrAgentEdit(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    defer r.Body.Close()
    agentID := opt.Params["agent"]

    dfc := s.GetDataSource("persistence")

    p, ok := dfc.(persistence.IPersistance);
    if !ok {
        utl.Log("Invalid store ")
        return
    }
    c := p.GetCollection("agentConfigs")

    var agentConfig agent.AgentConfig

    links := map[string]map[string]string{
        "self": map[string]string {
            "href": "api/agentsConfig/" + agentID,
        },
    }
    embedded := map[string]string{}

    switch r.Method {
    case "POST":
        decoder := json.NewDecoder(r.Body)
        decoder.Decode(&agentConfig)
        agentConfig.AgentID = agentID
        _, err := c.Upsert(
            bson.M{"agentId": agentID},
            agentConfig,
        )
        if err != nil {
            utl.Log(err)
        }
        json.NewEncoder(w).Encode(helpers.ServeHal(agentConfig, embedded, links))

    case "GET":
        err := c.Find(bson.M{
            "agentId": agentID,
        }).One(&agentConfig)

        if err != nil {
            utl.Log("AgentConfig not found err=", err)
            return
        }

        links := map[string]map[string]string{
            "self": map[string]string {
                "href": "api/agentsConfig/" + agentID,
            },
        }
        json.NewEncoder(w).Encode(helpers.ServeHal(agentConfig, embedded, links))

    default:
    }
}
