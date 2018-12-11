package controllers

import (
	"net/http"
    "github.com/smart-evolution/smarthome/processes/webserver/controllers/utils"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
    "github.com/coda-it/gowebserver/store"
)

// CtrDashboard - controller for agents list
func CtrDashboard(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    utils.RenderTemplate(w, r, "dashboard", sm, s)
}
