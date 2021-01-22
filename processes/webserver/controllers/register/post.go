package register

import (
	"github.com/coda-it/goutils/hash"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	userModel "github.com/smart-evolution/shapi/domain/models/user"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// CtrRegisterPost - post handler
func (c *Controller) CtrRegisterPost(w http.ResponseWriter, r *http.Request, opt router.URLOptions, sm session.ISessionManager, s store.IStore) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	newUser := userModel.User{
		ID:       bson.NewObjectId(),
		Username: username,
		Password: password,
		APIKey:   hash.EncryptString(username + password),
	}

	err := c.UserUsecases.RegisterUser(newUser)

	if err != nil {
		logger.Log("error while registering user", err)
		handlers.HandleError(w, "/register", "user registration failed", http.StatusInternalServerError)
		return
	}

	logger.Log("registered user '" + newUser.Username + "'")
}
