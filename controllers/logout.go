package controllers

import (
    "net/http"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
)

// AuthenticateLogout - logout user
func AuthenticateLogout(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    utils.ClearSession(w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
