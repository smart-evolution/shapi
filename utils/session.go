package utils

import (
    "net/http"
)

func GetSessionId(r *http.Request) string {
    sessionCookie, err := r.Cookie("sid")

    if err != nil {
        return ""
    }

    return sessionCookie.Value
}
