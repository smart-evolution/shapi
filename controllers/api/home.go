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
        Command:    "SELECT * FROM home ORDER BY time DESC LIMIT 30",
        Database:   "smarthome",
    }

    resp, err := services.InfluxClient.Query(q)

    if err != nil {
        log.Println(err)
    }

    res := resp.Results[0].Series[0]

    var (
        times           []string
        temperatures    []string
        presences       []string
    )

    for _, serie := range res.Values {
        times = append(times, serie[0].(string))
        temperatures = append(temperatures, serie[3].(string))
        presences = append(presences, serie[2].(string))
    }

    data := struct {
        Time        []string  `json:"time"`
		Temperature []string  `json:"temperature"`
        Presence    []string  `json:"presence"`
	} {
        times,
        temperatures,
        presences,
	}

	json.NewEncoder(w).Encode(data)
}

