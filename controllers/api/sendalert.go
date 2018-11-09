package api

import (
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
    "github.com/smart-evolution/smarthome/utils"
)

// CtrSendAlert - api controller for sending alerts to agents
func CtrSendAlert(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    if r.Method == "POST" {
        utils.SendAlert = true
    }

    data := struct {
        SendAlert    string  `json:"isAlerts"`
    } {
        strconv.FormatBool(utils.SendAlert),
    }

    json.NewEncoder(w).Encode(data)
}

