package api

import (
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/smart-evolution/smarthome/processes/webserver/controllers/api/agents"
    "github.com/smart-evolution/smarthome/datasources/dataflux"
    "github.com/smart-evolution/smarthome/datasources/state"
    "github.com/coda-it/gowebserver/store"
    "github.com/coda-it/gowebserver/helpers"
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
            utils.Log("Store should implement IDataFlux")
            return
        }
        st := s.GetDataSource("state")

        state, ok := st.(state.IState);
        if !ok {
            utils.Log("Store should implement IState")
            return
        }

        agentsType1, err := agents.FetchType1(agentID, df)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            utils.Log(err)
            return
        }

        agentsList = append(agentsList, agentsType1...)

        agentsType2, err := agents.FetchType2(agentID, state.Agents())
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            utils.Log(err)
            return
        }

        agentsList = append(agentsList, agentsType2...)

        if len(agentsList) == 0 {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        data := map[string]string{
            "count": strconv.Itoa(len(agentsList)),
        }

        links := map[string]map[string]string{
            "self": map[string]string {
                "href": "/api/agents/" + agentID,
            },
        }

        embedded := map[string]interface{}{
            "agents": agentsList,
        }

        json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))

    case "POST":
        dfc := s.GetDataSource("state")
        st, ok := dfc.(state.IState);
        if !ok {
            utils.Log("Store should implement IState")
            return
        }

        agent, err := st.AgentByID(agentID)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            utils.Log(err)
            return
        }

        _, err = http.Get(agent.URL())

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            utils.Log(err)
            return
        }

    default:
    }
}

