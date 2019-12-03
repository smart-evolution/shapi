package api

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

// CtrFront - api landing page
func CtrFront(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	data := struct{}{}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/",
		},
	}

	embedded := map[string]string{}

	switch r.Method {
	case "GET":
		handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
		return
	}
}
