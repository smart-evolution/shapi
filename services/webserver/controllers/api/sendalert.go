package api

import (
    "log"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/smart-evolution/smarthome/state"
    "github.com/coda-it/gowebserver/store"
)

// CtrSendAlert - api controller for sending alerts to agents
func CtrSendAlert(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    dfc := s.GetDataSource("state")

    st, ok := dfc.(state.IState);
    if !ok {
        log.Println("controllers: Invalid store ")
        return
    }


    if r.Method == "POST" {
        st.SetSendAlert(true)
    }

    data := struct {
        SendAlert    string  `json:"isAlerts"`
    } {
        strconv.FormatBool(st.SendAlert()),
    }

    json.NewEncoder(w).Encode(data)
}

