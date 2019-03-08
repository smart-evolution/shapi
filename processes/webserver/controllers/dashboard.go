package controllers

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/smarthome/processes/webserver/controllers/utils"
	"net/http"
)

// CtrDashboard - controller for agents list
func CtrDashboard(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	params := make(map[string]interface{})
	utils.RenderTemplate(w, r, "dashboard", sm, s, params)
}
