package api

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"net/http"
)

// CtrNotFound - api 404 page
func CtrNotFound(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	data := struct {
		Message string `json:"message"`
	}{
		"Endpoint not found",
	}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": "/api/404",
		},
	}

	embedded := map[string]string{}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
		return
	}
}
