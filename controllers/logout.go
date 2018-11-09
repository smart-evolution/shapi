package controllers

import (
    "net/http"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
)

// AuthenticateLogout - logout user
func AuthenticateLogout(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    utils.ClearSession(w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
