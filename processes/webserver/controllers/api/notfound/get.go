package notfound

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

// CtrNotFound - api 404 page
func (c *Controller) CtrNotFound(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)
	handlers.HandleError(w, "/api/404", "endpoint not found", http.StatusNotFound)
}
