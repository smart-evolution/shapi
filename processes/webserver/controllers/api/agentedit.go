package api

import (
    "net/http"
    "github.com/smart-evolution/smarthome/datasources/persistence"
    utl "github.com/smart-evolution/smarthome/utils"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/coda-it/gowebserver/store"
)

type AgentConfig struct {
    ID          bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
    AgentID     string 	        `json:"agentId" bson:"agentId,omitempty"`
    TmpAdjust   string          `json:"temperature" bson:"tmpAdjustment"`
}

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

    var agentConfig AgentConfig

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

    case "GET":
        err := c.Find(bson.M{
            "agentID": agentID,
        }).One(&agentConfig)

        if err != nil {
            utl.Log("AgentConfig not found err=", err)
            return
        }

        utl.Log("webserver/authenticateUser: Logged in as user", agentConfig)

    default:
    }
}
