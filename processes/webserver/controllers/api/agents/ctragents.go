package agents

import (
	"encoding/base64"
	"github.com/coda-it/goutils/logger"
	goutilsSession "github.com/coda-it/goutils/session"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/constants"
	"github.com/smart-evolution/shapi/datasources"
	"github.com/smart-evolution/shapi/datasources/dataflux"
	"github.com/smart-evolution/shapi/datasources/persistence"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/models/type1"
	ctrHandlers "github.com/smart-evolution/shapi/processes/webserver/controllers/api/agents/handlers"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	webSrvUtils "github.com/smart-evolution/shapi/processes/webserver/utils"
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
		dsp := s.GetDataSource(datasources.Persistence)
		p, ok := dsp.(persistence.IPersistance)

		if !ok {
			logger.Log("store should implement persistence")
			handlers.HandleError(w, href, "controller store error", http.StatusInternalServerError)
			return
		}

		sessionID, _ := goutilsSession.GetSessionID(r, constants.SessionKey)
		isLogged := sm.IsExist(sessionID)

		if !isLogged {
			authorization := r.Header.Get("Authorization")

			if authorization != "" {
				authData := strings.Split(authorization, " ")
				token := authData[1]
				credentials, err := base64.StdEncoding.DecodeString(token)

				if err != nil {
					logger.Log("Decoding auth token failed")
				}

				credArr := strings.Split(string(credentials), ":")
				username := credArr[0]
				password := credArr[1]

				webSrvUtils.CreateClientSession(w, r, username, password, p, sm)
			}
		}

		var list []AgentJSON
		dfc := s.GetDataSource(datasources.Dataflux)

		df, ok := dfc.(dataflux.IDataFlux)
		if !ok {
			logger.Log("store should implement dataflux")
			handlers.HandleError(w, href, "controller store error", http.StatusInternalServerError)
			return
		}
		st := s.GetDataSource(datasources.State)

		is, ok := st.(state.IState)
		if !ok {
			logger.Log("store should implement state")
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
					logger.Log(err)
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
					logger.Log(err)
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
				logger.Log("type assertion error")
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
		dss := s.GetDataSource(datasources.State)
		st, ok := dss.(state.IState)
		if !ok {
			logger.Log("store should implement state")
			handlers.HandleError(w, href, "controller store error", http.StatusInternalServerError)
			return
		}

		ia, err := st.AgentByID(agentID)

		if err != nil {
			logger.Log("agent with id = " + agentID + " not found")
			handlers.HandleError(w, href, "agent not found", http.StatusNotFound)
			return
		}

		foundAgent, ok := ia.(*agent.Agent)

		if !ok {
			logger.Log("type assertion error")
			return
		}

		_, err = http.Get(foundAgent.IP)

		if err != nil {
			logger.Log("requesting agent with IP = " + foundAgent.IP + " failed")
			handlers.HandleError(w, href, "error contacting agent", http.StatusInternalServerError)
			return
		}

	case "DELETE":
		dss := s.GetDataSource(datasources.State)
		st := dss.(state.IState)
		ctrHandlers.Delete(st, w, href, agentID)

	default:
	}
}
