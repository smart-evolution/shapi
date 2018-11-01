package controllers

import (
    "net/http"
    "time"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
)

func AuthenticateLogout(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    cookie := http.Cookie {
        Path: "/",
        Name: "sid",
        Expires: time.Now().Add(-100 * time.Hour),
        MaxAge: -1 }

    http.SetCookie(w, &cookie)

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
