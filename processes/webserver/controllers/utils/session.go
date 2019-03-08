package utils

import (
	"net/http"
	"time"
)

// CreateSessionID - creates a new session ID
func CreateSessionID(user string, pass string, time string) string {
	return HashString(user + pass + time)
}

// GetSessionID - get user session ID
func GetSessionID(r *http.Request) (string, error) {
	sessionCookie, err := r.Cookie("sid")

	if err != nil {
		return "", err
	}

	return sessionCookie.Value, nil
}

// ClearSession - remove session cookie
func ClearSession(w http.ResponseWriter) {
	cookie := http.Cookie{
		Path:    "/",
		Name:    "sid",
		Expires: time.Now().Add(-100 * time.Hour),
		MaxAge:  -1}
	http.SetCookie(w, &cookie)
}
