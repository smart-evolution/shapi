package front

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

// CtrFrontAll - api landing page
func (c *Controller) CtrFrontAll(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	data := struct{}{}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/",
		},
	}

	embedded := map[string]string{}

	handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
}
