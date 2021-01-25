package agents

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"net/http"
)

// CtrAgentsOptions - options handler
func (c *Controller) CtrAgentsOptions(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	c.CorsHeaders(w, r)

	switch r.Method {
	case "OPTIONS":
		return
	}
}
