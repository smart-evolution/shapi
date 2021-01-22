package sendalert

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
	"strconv"
)

const sendAlertHref string = "/api/sendalert"

// CtrSendAlertAll - api controller for sending alerts to agents
func (c *Controller) CtrSendAlertAll(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	if r.Method == "POST" {
		c.UserUsecases.SetSendAlert(true)
	}

	data := struct {
		SendAlert string `json:"isAlerts"`
	}{
		strconv.FormatBool(c.UserUsecases.SendAlert()),
	}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": sendAlertHref,
		},
	}

	embedded := map[string]string{}

	handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
}
