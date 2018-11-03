package utils

import (
    "net/http"
)

// GetSessionID - get user session ID
func GetSessionID(r *http.Request) (string, error) {
    sessionCookie, err := r.Cookie("sid")

    if err != nil {
        return nil, err
    }

    return sessionCookie.Value, nil
}
