package agentconfigs

import (
	"encoding/base64"
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
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
	defer r.Body.Close()
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	agentID := opt.Params["agent"]
	href := "api/agentsConfig/" + agentID

	dfc := s.GetDataSource(datasources.Persistence)

	p, ok := dfc.(persistence.IPersistance)
	if !ok {
		utl.Log("Invalid store - should implement `IPersistance`")
		handlers.HandleError(w, href, "controller store error", http.StatusInternalServerError)
		return
	}
	c := p.GetCollection("agentConfigs")

	var config agent.Config

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

		var list []agent.Config

		err := c.Find(bson.M{}).All(&list)

		if err != nil {
			utl.Log("AgentConfig not found")
			handlers.HandleError(w, href, "agentConfig not found", http.StatusNotFound)
			return
		}

		data := map[string]string{
			"count": strconv.Itoa(len(list)),
		}

		embedded := map[string]interface{}{
			"configs": list,
		}
		json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))

	case "POST":
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&config)
		config.AgentID = agentID
		_, err := c.Upsert(
			bson.M{"agentId": agentID},
			config,
		)

		if err != nil {
			utl.Log("error while posting agentConfig")
			handlers.HandleError(w, href, "error while posting agentConfig", http.StatusInternalServerError)
			return
		}

		embedded := map[string]string{}
		json.NewEncoder(w).Encode(helpers.ServeHal(config, embedded, links))

	default:
	}
}
