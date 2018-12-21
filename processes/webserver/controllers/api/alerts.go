package api

import (
    "log"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/smart-evolution/smarthome/datasources/state"
    "github.com/coda-it/gowebserver/store"
    "github.com/coda-it/gowebserver/helpers"
)

// CtrAlerts - api controller for sending alerts to agents
func CtrAlerts(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    dfc := s.GetDataSource("state")

    st, ok := dfc.(state.IState);
    if !ok {
        log.Println("webserver/CtrAlerts: Invalid store ")
        return
    }

    if r.Method == "POST" {
        st.SetIsAlerts(!st.IsAlerts())
    }

    data := struct {
        IsAlerts    string  `json:"isAlerts"`
    }{
        strconv.FormatBool(st.IsAlerts()),
    }

    links := map[string]map[string]string{
        "self": map[string]string {
            "href": "/api/alerts",
        },
    }

    embedded := map[string]string{}

    json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
}

