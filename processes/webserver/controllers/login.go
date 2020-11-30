package controllers

import (
	"github.com/coda-it/goutils/hash"
	"github.com/coda-it/goutils/logger"
	goutilsSession "github.com/coda-it/goutils/session"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/constants"
	"github.com/smart-evolution/shapi/datasources/persistence"
	"github.com/smart-evolution/shapi/processes/webserver/utils"
	"net/http"
)

// Authenticate - handle login page and login process
func Authenticate(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	defer r.Body.Close()

	switch r.Method {
	case "POST":
		sessionID, _ := goutilsSession.GetSessionID(r, constants.SessionKey)
		isLogged := sm.IsExist(sessionID)

		if !isLogged {
			username := r.PostFormValue("username")
			password := hash.EncryptString(r.PostFormValue("password"))

			dfc := s.GetDataSource("persistence")

			p, ok := dfc.(persistence.IPersistance)
			if !ok {
				logger.Log("Invalid store")
				return
			}

			isSession := utils.CreateClientSession(w, r, username, password, p, sm)

			if isSession {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/login?err", http.StatusSeeOther)
			}
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	default:
	}
}
