package controllers

import (
	"net/http"
    "github.com/smart-evolution/smarthome/controllers/utils"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
)

// CtrDashboard - controller for agents list
func CtrDashboard(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    utils.RenderTemplate(w, r, "dashboard", sm)
}
