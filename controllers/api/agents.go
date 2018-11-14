package api

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
    "github.com/smart-evolution/smarthome/controllers/api/agents"
    "github.com/smart-evolution/smarthome/models"
)

// CtrAgents - controller for retrieving agents list data
func CtrAgents(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    agentID := opt.Params["agent"]

    switch r.Method {
    case "GET":
        var agentsList []agents.AgentJSON

        agentsType1, err := agents.FetchType1(agentID)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println("controllers: ", err)
            return
        }

        agentsList = append(agentsList, agentsType1...)

        agentsType2, err := agents.FetchType2(agentID)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println("controllers: ", err)
            return
        }

        agentsList = append(agentsList, agentsType2...)

        if len(agentsList) == 0 {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        json.NewEncoder(w).Encode(agentsList)

    case "POST":
        agent, err := models.FindAgentByID(agentID)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println("controllers: ", err)
            return
        }

        _, err = http.Get(agent.URL)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println("controllers: ", err)
            return
        }

    default:
    }
}

