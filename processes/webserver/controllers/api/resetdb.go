package api

import (
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

const resetHref string = "/api/reset"

// CtrResetDb - resets state
func CtrResetDb(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	switch r.Method {
	case "POST":
		data := struct {
			Message string `json:"message"`
		}{
			"state cleared",
		}

		st := s.GetDataSource(datasources.State)

		is, ok := st.(state.IState)
		if !ok {
			logger.Log("store should implement state")
			handlers.HandleError(w, resetHref, "controller store error", http.StatusInternalServerError)
			return
		}

		is.Reset()

		links := map[string]map[string]string{
			"self": map[string]string{
				"href": resetHref,
			},
		}

		embedded := map[string]string{}

		handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
		return
	}
}
