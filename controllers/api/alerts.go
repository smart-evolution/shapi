package api

import (
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
    "github.com/oskarszura/smarthome/utils"
)

// CtrAlerts - api controller for sending alerts to agents
func CtrAlerts(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    if r.Method == "POST" {
        utils.IsAlerts = !utils.IsAlerts
    }

    data := struct {
        IsAlerts    string  `json:"isAlerts"`
    } {
        strconv.FormatBool(utils.IsAlerts),
    }

    json.NewEncoder(w).Encode(data)
}

