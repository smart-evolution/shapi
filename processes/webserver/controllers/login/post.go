package login

import (
	"github.com/coda-it/goutils/hash"
	goutilsSession "github.com/coda-it/goutils/session"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/constants"
	"net/http"
)

// CtrLoginPost - handle login page and login process
func (c *Controller) CtrLoginPost(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	defer r.Body.Close()

	sessionID, _ := goutilsSession.GetSessionID(r, constants.SessionKey)
	isLogged := sm.IsExist(sessionID)

	if !isLogged {
		username := r.PostFormValue("username")
		password := hash.EncryptString(r.PostFormValue("password"))

		isSession := c.UserUsecases.CreateClientSession(w, r, username, password, sm)

		if isSession {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login?err", http.StatusSeeOther)
		}
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
