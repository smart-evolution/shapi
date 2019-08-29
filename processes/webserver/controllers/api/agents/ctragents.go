package agents

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/smarthome/datasources"
	"github.com/smart-evolution/smarthome/datasources/dataflux"
	"github.com/smart-evolution/smarthome/datasources/state"
	"github.com/smart-evolution/smarthome/models/agent/types"
	"github.com/smart-evolution/smarthome/utils"
	"net/http"
	"strconv"
)

// CtrAgents - controller for retrieving agents list data
func CtrAgents(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	agentID := opt.Params["agent"]
	period := r.URL.Query().Get("period")

	if period == "" {
		period = "30"
	}

	switch r.Method {
	case "GET":
		var list []AgentJSON
		dfc := s.GetDataSource(datasources.Dataflux)

		df, ok := dfc.(dataflux.IDataFlux)
		if !ok {
			utils.Log("Store should implement IDataFlux")
			return
		}
		st := s.GetDataSource(datasources.State)

		state, ok := st.(state.IState)
		if !ok {
			utils.Log("Store should implement IState")
			return
		}

		cnfAgents := state.Agents()
		for _, a := range cnfAgents {
			var (
				data interface{}
				err  error
			)
			rawType := a.RawType()

			if rawType == types.Type1 {
				data, err = FetchType1Data(a.ID(), period, df)
				if err != nil {
					utils.Log(err)
				}
			} else if rawType == types.Type2 {
				data, err = FetchType2(a.ID(), state.Agents())

				if err != nil {
					utils.Log(err)
				}
			} else if rawType == types.Jeep {
				data = nil
			} else {
				data = nil
			}

			agentJSON := AgentJSON{
				ID:        a.ID(),
				Name:      a.Name(),
				Data:      data,
				AgentType: a.AgentType(),
				IP:        a.IP(),
				IsOnline:  a.IsOnline(),
			}
			list = append(list, agentJSON)
		}

		if len(list) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		data := map[string]string{
			"count": strconv.Itoa(len(list)),
		}

		links := map[string]map[string]string{
			"self": map[string]string{
				"href": "/api/agents/" + agentID,
			},
		}

		embedded := map[string]interface{}{
			"agents": list,
		}

		json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))

	case "POST":
		dfc := s.GetDataSource("state")
		st, ok := dfc.(state.IState)
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

		_, err = http.Get(agent.IP())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.Log(err)
			return
		}

	default:
	}
}
