package resetdb

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"net/http"
)

const resetHref string = "/api/reset"

// CtrResetDbPost - resets state
func (c *Controller) CtrResetDbPost(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	c.CorsHeaders(w, r)

	switch r.Method {
	case "POST":
		data := struct {
			Message string `json:"message"`
		}{
			"state cleared",
		}

		err := c.PlatformUsecases.Drop()

		if err != nil {
			// throw error
		}

		links := map[string]map[string]string{
			"self": map[string]string{
				"href": resetHref,
			},
		}

		embedded := map[string]string{}

		c.HandleJSONResponse(w, data, embedded, links, http.StatusOK)
		return
	}
}
