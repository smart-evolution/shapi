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
        time            string
        temperature     string
        presence        string
    )

    for _, serie := range res.Values {
        if serie[0] != nil {
            time = serie[0].(string)
        } else {
            time = ""
        }
        if serie[3] != nil {
            temperature = serie[3].(string)
        } else {
            temperature = ""
        }
        if serie[2] != nil {
            presence = serie[2].(string)
        } else {
            presence = ""
        }

        times = append(times, time)
        temperatures = append(temperatures, temperature)
        presences = append(presences, presence)
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

