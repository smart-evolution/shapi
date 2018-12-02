package controllers

import (
    "net/http"
    "github.com/smart-evolution/smarthome/services/webserver/controllers/utils"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/coda-it/gowebserver/store"
)

// AuthenticateLogout - logout user
func AuthenticateLogout(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    utils.ClearSession(w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
