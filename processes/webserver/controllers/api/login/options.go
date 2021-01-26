package login

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"net/http"
)

// CtrLoginOptions - options handler
func (c *Controller) CtrLoginOptions(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	defer r.Body.Close()
	c.CorsHeaders(w, r)
}
