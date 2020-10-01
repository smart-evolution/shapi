package utils

import (
	"errors"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/gowebserver/session"
	"github.com/smart-evolution/shapi/datasources/persistence"
	"github.com/smart-evolution/shapi/models/user"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

// AuthenticateByCredentials - authenticate user with credentials
func AuthenticateByCredentials(username string, password string, p persistence.IPersistance) (user.User, error) {
	usr, err := p.FindOneUser(bson.M{
		"username": username,
		"password": password,
	})

	if err != nil {
		msg := "user not found"
		logger.Log(msg)
		return user.User{}, errors.New(msg)
	}

	logger.Log("logged in as user", usr)

	return usr, nil
}

// CreateClientSession - authenticate uer with credentials and create session cookie
func CreateClientSession(w http.ResponseWriter, r *http.Request, username string, password string, p persistence.IPersistance, sm session.ISessionManager) bool {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	authenticatedUser, err := AuthenticateByCredentials(username, password, p)

	if err == nil {
		t := time.Now()
		timeStr := t.Format(time.RFC850)
		cookieValue := CreateSessionID(username, password, timeStr)

		cookie := http.Cookie{
			Name:    "sid",
			Value:   cookieValue,
			Expires: expiration}

		session := sm.Create(cookieValue)
		session.Set("user", authenticatedUser)

		http.SetCookie(w, &cookie)
		return true
	}
	return false
}
