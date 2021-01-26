package notfound

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"net/http"
)

// CtrNotFound - api 404 page
func (c *Controller) CtrNotFound(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	defer r.Body.Close()
	c.HandleError(w, "/api/404", "endpoint not found", http.StatusNotFound)
}
