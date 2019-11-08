package controllers

import (
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/smarthome/datasources/persistence"
	"github.com/smart-evolution/smarthome/models/user"
	"github.com/smart-evolution/smarthome/processes/webserver/controllers/utils"
	utl "github.com/smart-evolution/smarthome/utils"
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
			utl.Log("Invalid store")
			return
		}

		c := p.GetCollection("users")

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		newUser = &user.User{
			ID:       bson.NewObjectId(),
			Username: username,
			Password: password,
			APIKey:   utils.HashString(username + password),
		}

		err := c.Insert(newUser)
		if err != nil {
			utl.Log("Registering user '" + newUser.Username + "' failed")
			return
		}

		utl.Log("Registered user '" + newUser.Username + "'")
	default:
	}
}
