package alerts

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"net/http"
	"strconv"
)

// CtrAlertsPost - post handler
func (c *Controller) CtrAlertsPost(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	c.CorsHeaders(w, r)

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/alerts",
		},
	}

	embedded := map[string]string{}

	data := struct {
		IsAlerts string `json:"isAlerts"`
	}{
		strconv.FormatBool(c.UserUsecases.IsAlerts()),
	}
	c.HandleJSONResponse(w, data, embedded, links, http.StatusOK)
}