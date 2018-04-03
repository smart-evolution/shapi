package api

import (
    "log"
    "net/http"
    "encoding/json"
    services "github.com/oskarszura/smarthome/services"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
    "github.com/influxdata/influxdb/client/v2"
)

func CtrHome(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    q := client.Query{
        Command:    "SELECT * FROM home",
        Database:   "smarthome",
    }

    resp, err := services.InfluxClient.Query(q)

    if err != nil {
        log.Println(err)
    }

    res := resp.Results[0].Series[0]

    data := struct {
		Temperature string  `json:"temperature"`
        Presence    string  `json:"presence"`
	} {
        res.Values[0][3].(string),
        res.Values[0][2].(string),
	}

	json.NewEncoder(w).Encode(data)
}

