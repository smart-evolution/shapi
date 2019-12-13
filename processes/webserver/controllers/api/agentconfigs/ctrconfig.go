package agentconfigs

import (
	"encoding/base64"
	"encoding/json"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources"
	"github.com/smart-evolution/shapi/datasources/persistence"
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"github.com/smart-evolution/shapi/processes/webserver/utils"
	utl "github.com/smart-evolution/shapi/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)

// CtrAgentConfig - controller for agents list
func CtrAgentConfig(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	agentID := opt.Params["agent"]
	href := "api/agentsConfig/" + agentID

	dfc := s.GetDataSource(datasources.Persistence)

	p, ok := dfc.(persistence.IPersistance)
	if !ok {
		utl.Log("Invalid store - should implement `IPersistance`")
		handlers.HandleError(w, href, "controller store error", http.StatusInternalServerError)
		return
	}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": href,
		},
	}

	switch r.Method {
	case "OPTIONS":
		return
	case "GET":
		sessionID, _ := utils.GetSessionID(r)
		isLogged := sm.IsExist(sessionID)

		if !isLogged {
			authorization := r.Header.Get("Authorization")

			if authorization != "" {
				authData := strings.Split(authorization, " ")
				token := authData[1]
				credentials, err := base64.StdEncoding.DecodeString(token)

				if err != nil {
					utl.Log("Encoding credentials failed")
					handlers.HandleError(w, href, "error encoding credentials", http.StatusInternalServerError)
					return
				}

				credArr := strings.Split(string(credentials), ":")
				username := credArr[0]
				password := credArr[1]

				utils.CreateClientSession(w, r, username, password, p, sm)
			}
		}

		list, err := p.FindAllAgentConfigs(bson.M{})

		if err != nil {
			msg := "AgentConfig not found"
			utl.Log(msg)
			handlers.HandleError(w, href, msg, http.StatusNotFound)
			return
		}

		data := map[string]string{
			"count": strconv.Itoa(len(list)),
		}

		embedded := map[string]interface{}{
			"configs": list,
		}
		handlers.HandleResponse(w, data, embedded, links, http.StatusOK)

	case "POST":
		var config agent.Config
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		decoder.Decode(&config)
		config.AgentID = agentID

		err := p.Upsert("agentConfigs", bson.M{"agentId": agentID}, config)

		if err != nil {
			utl.Log("error while posting agentConfig")
			handlers.HandleError(w, href, "error while posting agentConfig", http.StatusInternalServerError)
			return
		}

		embedded := map[string]string{}
		handlers.HandleResponse(w, config, embedded, links, http.StatusOK)

	default:
	}
}
