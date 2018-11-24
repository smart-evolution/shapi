package utils

import (
    "fmt"
    "time"
    "net/http"
    "crypto/sha1"
)

// CreateSessionID - creates a new session ID
func CreateSessionID() string {
    t := time.Now()
    val := []byte(t.Format(time.RFC850))
    h := sha1.New()
    h.Write(val)

    return fmt.Sprintf("%x", h.Sum(nil))
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
    cookie := http.Cookie {
        Path: "/",
        Name: "sid",
        Expires: time.Now().Add(-100 * time.Hour),
        MaxAge: -1 }

    http.SetCookie(w, &cookie)
}
