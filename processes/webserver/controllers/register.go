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
	case "GET":
		params := make(map[string]interface{})
		utils.RenderTemplate(w, r, "register", sm, s, params)

	case "POST":
		var newUser *user.User

		dfc := s.GetDataSource("persistence")

		p, ok := dfc.(persistence.IPersistance)
		if !ok {
			utl.Log("Invalid store")
			return
		}

		c := p.GetCollection("users")

		newUser = &user.User{
			ID:       bson.NewObjectId(),
			Username: r.PostFormValue("username"),
			Password: utils.HashString(r.PostFormValue("password")),
		}

		err := c.Insert(newUser)
		if err != nil {
			utl.Log(err)
		}

		utl.Log("Registered user", newUser)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
	}
}
