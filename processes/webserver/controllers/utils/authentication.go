package utils

import (
	"errors"
	"github.com/coda-it/gowebserver/session"
	"github.com/smart-evolution/smarthome/datasources/persistence"
	"github.com/smart-evolution/smarthome/models/user"
	utl "github.com/smart-evolution/smarthome/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func AuthenticateByCredentials(username string, password string, p persistence.IPersistance) (user.User, error) {
	var user user.User

	c := p.GetCollection("users")

	err := c.Find(bson.M{
		"username": username,
		"password": password,
	}).One(&user)

	if err != nil {
		utl.Log("User not found err=", err)
		return user, errors.New("User not found")
	}

	utl.Log("webserver/authenticateUser: Logged in as user", user)

	return user, nil
}

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
	} else {
		return false
	}
}
