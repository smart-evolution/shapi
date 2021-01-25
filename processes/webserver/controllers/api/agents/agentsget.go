package agents

import (
	"encoding/base64"
	"github.com/coda-it/goutils/logger"
	goutilsSession "github.com/coda-it/goutils/session"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/constants"
	"github.com/smart-evolution/shapi/domain/models/agent"
	"github.com/smart-evolution/shapi/domain/models/type1"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
	"strconv"
	"strings"
)

// CtrAgentsGet - get handler
func (c *Controller) CtrAgentsGet(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	agentID := opt.Params["agent"]
	period := r.URL.Query().Get("period")

	href := "/api/agents/" + agentID

	if period == "" {
		period = "30"
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

			c.UserUsecases.CreateClientSession(w, r, username, password, sm)
		}
	}

	var list []AgentJSON

	for _, ia := range c.AgentUsecases.Agents() {
		var (
			data interface{}
			err  error
		)

		switch a := ia.(type) {
		case *agent.Agent:
			data, err = c.AgentUsecases.FetchType1Data(a.ID, period)
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
			data, err = c.AgentUsecases.FetchType1Data(a.ID, period)
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

	handlers.HandleJSONResponse(w, data, embedded, links, http.StatusOK)
}
