package resetdb

import (
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"net/http"
)

const resetHref string = "/api/reset"

// CtrResetDbPost - resets state
func (c *Controller) CtrResetDbPost(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	defer r.Body.Close()

	err := c.PlatformUsecases.Drop()
	if err != nil {
		if err != nil {
			logger.Log("resetdb errored")
			c.HandleError(w, "/api/reset", "resetdb errored", http.StatusInternalServerError)
			return
		}
	}

	data := struct {
		Message string `json:"message"`
	}{
		"state cleared",
	}
	links := map[string]map[string]string{
		"self": map[string]string{
			"href": resetHref,
		},
	}
	embedded := map[string]string{}

	c.HandleJSONResponse(w, data, embedded, links, http.StatusOK)
}
