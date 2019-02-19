package controllers

import (
	"net/http"
	"github.com/smart-evolution/smarthome/processes/webserver/controllers/utils"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
    "github.com/coda-it/gowebserver/store"
)

// NotFound - controller for 404 requests
func NotFound(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    params := make(map[string]interface{})
    utils.RenderTemplate(w, r, "404", sm, s, params)
}
