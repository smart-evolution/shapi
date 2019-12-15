package controllers

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources/persistence"
	"github.com/smart-evolution/shapi/models/user"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"github.com/smart-evolution/shapi/processes/webserver/utils"
	utl "github.com/smart-evolution/shapi/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// Register - handle register page and register user process
func Register(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	switch r.Method {
	case "POST":
		var newUser *user.User
		dfc := s.GetDataSource("persistence")
		p, ok := dfc.(persistence.IPersistance)

		if !ok {
			handlers.HandleError(w, "/register", "controller store error", http.StatusInternalServerError)
			return
		}

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		newUser = &user.User{
			ID:       bson.NewObjectId(),
			Username: username,
			Password: password,
			APIKey:   utils.HashString(username + password),
		}

		err := p.Insert("users", newUser)

		if err != nil {
			utl.Log("error while registering user", err)
			handlers.HandleError(w, "/register", "user registration failed", http.StatusInternalServerError)
			return
		}

		utl.Log("registered user '" + newUser.Username + "'")
	default:
	}
}
