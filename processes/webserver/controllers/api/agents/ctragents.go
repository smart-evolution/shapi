package agents

import (
	"encoding/base64"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources"
	"github.com/smart-evolution/shapi/datasources/dataflux"
	"github.com/smart-evolution/shapi/datasources/persistence"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/models/type1"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"github.com/smart-evolution/shapi/processes/webserver/utils"
	utl "github.com/smart-evolution/shapi/utils"
	"net/http"
	"strconv"
	"strings"
)

// CtrAgents - controller for retrieving agents list data
func CtrAgents(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	agentID := opt.Params["agent"]
	period := r.URL.Query().Get("period")

	href := "/api/agents/" + agentID

	if period == "" {
		period = "30"
	}

	switch r.Method {
	case "OPTIONS":
		return
	case "GET":
		pc := s.GetDataSource("persistence")

		p, ok := pc.(persistence.IPersistance)
		if !ok {
			utl.Log("Invalid store")
			return
		}

		sessionID, _ := utils.GetSessionID(r)
		isLogged := sm.IsExist(sessionID)

		if !isLogged {
			authorization := r.Header.Get("Authorization")

			if authorization != "" {
				authData := strings.Split(authorization, " ")
				token := authData[1]
				credentials, err := base64.StdEncoding.DecodeString(token)

				if err != nil {
					utl.Log("Decoding auth token failed")
				}

				credArr := strings.Split(string(credentials), ":")
				username := credArr[0]
				password := credArr[1]

				utils.CreateClientSession(w, r, username, password, p, sm)
			}
		}

		var list []AgentJSON
		dfc := s.GetDataSource(datasources.Dataflux)

		df, ok := dfc.(dataflux.IDataFlux)
		if !ok {
			utl.Log("Store should implement IDataFlux")
			handlers.HandleError(w, href, "controller store error", http.StatusInternalServerError)
			return
		}
		st := s.GetDataSource(datasources.State)

		is, ok := st.(state.IState)
		if !ok {
			utl.Log("store should implement IState")
			handlers.HandleError(w, href, "controller store error", http.StatusInternalServerError)
			return
		}

		for _, ia := range is.Agents() {
			var (
				data interface{}
				err  error
			)

			switch a := ia.(type) {
			case *agent.Agent:
				data, err = FetchType1Data(a.ID, period, df)
				if err != nil {
					utl.Log(err)
				}

				agentJSON := AgentJSON{
					ID:        a.ID,
					Name:      a.Name,
					Data:      data,
					AgentType: a.AgentType,
					IP:        a.IP,
					IsOnline:  a.IsOnline,
				}
				list = append(list, agentJSON)
			case *type1.Type1:
				data, err = FetchType1Data(a.ID, period, df)
				if err != nil {
					utl.Log(err)
				}

				agentJSON := AgentJSON{
					ID:        a.ID,
					Name:      a.Name,
					Data:      data,
					AgentType: a.AgentType,
					IP:        a.IP,
					IsOnline:  a.IsOnline,
				}
				list = append(list, agentJSON)
			default:
				utl.Log("type assertion error")
				continue
			}
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
				"href": href,
			},
		}

		embedded := map[string]interface{}{
			"agents": list,
		}

		handlers.HandleResponse(w, data, embedded, links, http.StatusOK)

	case "POST":
		dfc := s.GetDataSource("state")
		is, ok := dfc.(state.IState)
		if !ok {
			utl.Log("Store should implement IState")
			handlers.HandleError(w, href, "controller store error", http.StatusInternalServerError)
			return
		}

		ia, err := is.AgentByID(agentID)

		if err != nil {
			utl.Log("Agent with id = " + agentID + " not found")
			handlers.HandleError(w, href, "agent not found", http.StatusNotFound)
			return
		}

		a, ok := ia.(*agent.Agent)

		if !ok {
			utl.Log("type assertion error")
			return
		}

		_, err = http.Get(a.IP)

		if err != nil {
			utl.Log("Requesting agent with IP = " + a.IP + " failed")
			handlers.HandleError(w, href, "error contacting agent", http.StatusInternalServerError)
			return
		}

	default:
	}
}
