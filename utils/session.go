package utils

import (
    "net/http"
)

// GetSessionID - get user session ID
func GetSessionID(r *http.Request) string {
    sessionCookie, err := r.Cookie("sid")

    if err != nil {
        return ""
    }

    return sessionCookie.Value
}
