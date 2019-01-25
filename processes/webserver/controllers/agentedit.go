package controllers

import (
    "net/http"
    "github.com/smart-evolution/smarthome/datasources/persistence"
    utl "github.com/smart-evolution/smarthome/utils"
    "gopkg.in/mgo.v2/bson"
    "github.com/smart-evolution/smarthome/processes/webserver/controllers/utils"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/coda-it/gowebserver/store"
    "fmt"
)

type AgentConfig struct {
    ID          bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
    AgentID     bson.ObjectId 	`json:"agentId" bson:"agentId,omitempty"`
    TmpAdjust   string          `bson:"tmpAdjustment"`
}

// CtrAgentEdit - controller for agents list
func CtrAgentEdit(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    utils.RenderTemplate(w, r, "agentedit", sm, s)
    agentID := opt.Params["agent"]
    fmt.Println("------------------->", agentID)

    dfc := s.GetDataSource("persistence")

    p, ok := dfc.(persistence.IPersistance);
    if !ok {
        utl.Log("Invalid store ")
        return
    }

    var agentConfig AgentConfig

    c := p.GetCollection("agentConfigs")
    err := c.Find(bson.M{
        "agentID": agentID,
    }).One(&agentConfig)

    if err != nil {
        utl.Log("AgentConfig not found err=", err)
        return
    }

    utl.Log("webserver/authenticateUser: Logged in as user", agentConfig)
}
