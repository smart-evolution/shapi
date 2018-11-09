package controllers

import (
	"net/http"
	"github.com/smart-evolution/smarthome/controllers/utils"
	"github.com/oskarszura/gowebserver/router"
	"github.com/oskarszura/gowebserver/session"
)

// NotFound - controller for 404 requests
func NotFound(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    utils.RenderTemplate(w, r, "404", sm)
}
