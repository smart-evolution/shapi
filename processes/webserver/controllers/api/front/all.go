package front

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/utils"
	"net/http"
)

// CtrFrontAll - api landing page
func (c *Controller) CtrFrontAll(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	defer r.Body.Close()
	c.CorsHeaders(w, r)

	data := map[string]string{
		"version": utils.VERSION,
	}
	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/",
		},
	}
	embedded := map[string]string{}

	c.HandleJSONResponse(w, data, embedded, links, http.StatusOK)
}
