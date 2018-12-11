package api

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/smart-evolution/smarthome/services/webserver/controllers/api/agents"
    "github.com/smart-evolution/smarthome/datasources/dataflux"
    "github.com/smart-evolution/smarthome/datasources/state"
    "github.com/coda-it/gowebserver/store"
)

// CtrAgents - controller for retrieving agents list data
func CtrAgents(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    agentID := opt.Params["agent"]

    switch r.Method {
    case "GET":
        var agentsList []agents.AgentJSON
        dfc := s.GetDataSource("dataflux")

        df, ok := dfc.(dataflux.IDataFlux);
        if !ok {
            log.Println("controllers: Invalid store ")
            return
        }
        st := s.GetDataSource("state")

        state, ok := st.(state.IState);
        if !ok {
            log.Println("controllers: Invalid store ")
            return
        }

        agentsType1, err := agents.FetchType1(agentID, df)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println("controllers: ", err)
            return
        }

        agentsList = append(agentsList, agentsType1...)

        agentsType2, err := agents.FetchType2(agentID, state.Agents())
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
        dfc := s.GetDataSource("state")
        st, ok := dfc.(state.IState);
        if !ok {
            log.Println("controllers: Invalid store ")
            return
        }

        agent, err := st.FindAgentByID(agentID)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println("controllers: ", err)
            return
        }

        _, err = http.Get(agent.URL())

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            log.Println("controllers: ", err)
            return
        }

    default:
    }
}

