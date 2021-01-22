package agentconfigs

import (
	"encoding/base64"
	"github.com/coda-it/goutils/logger"
	goutilsSession "github.com/coda-it/goutils/session"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/constants"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)

// CtrAgentConfigGet -
func (c *Controller) CtrAgentConfigGet(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	agentID := opt.Params["agent"]
	href := "api/agentsConfig/" + agentID

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": href,
		},
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
				logger.Log("Encoding credentials failed")
				handlers.HandleError(w, href, "error encoding credentials", http.StatusInternalServerError)
				return
			}

			credArr := strings.Split(string(credentials), ":")
			username := credArr[0]
			password := credArr[1]

			c.UserUsecases.CreateClientSession(w, r, username, password, sm)
		}
	}

	list, err := c.AgentUsecases.FindAllAgentConfigs(bson.M{})

	if err != nil {
		msg := "AgentConfig not found"
		logger.Log(msg)
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
}
