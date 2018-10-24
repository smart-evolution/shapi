package controllers

import (
	"net/http"
    "github.com/oskarszura/smarthome/controllers/utils"
	"github.com/oskarszura/gowebserver/router"
	"github.com/oskarszura/gowebserver/session"
)

// CtrDashboard - controller for agents list
func CtrDashboard(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    utils.RenderTemplate(w, r, "dashboard", sm)
}
