package api

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

// CtrNotFound - api 404 page
func CtrNotFound(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	switch r.Method {
	case "GET":
		handlers.HandleError(w, "/api/404", "endpoint not found", http.StatusNotFound)
		return
	}
}
