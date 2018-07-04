package api

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/oskarszura/smarthome/services"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
    "github.com/influxdata/influxdb/client/v2"
)

type Agent struct {
    name    string
    data    AgentData
}

type AgentData struct {
    Time        []string  `json:"time"`
    Temperature []string  `json:"temperature"`
    Presence    []string  `json:"presence"`
    Gas         []string  `json:"gas"`
    Sound       []string  `json:"sound"`
}

func CtrHome(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    var data []Agent

    for _, agent := range services.Agents {
        q := client.Query{
            Command:    "SELECT time, temperature, presence, gas, sound, agent FROM home WHERE agent = '" + agent.Name + "' ORDER BY time DESC LIMIT 30",
            Database:   "smarthome",
        }

        resp, err := services.InfluxClient.Query(q)

        if err != nil {
            log.Println("services: ", err)
        }

        res := resp.Results[0].Series[0]

        var (
            times           []string
            temperatures    []string
            presences       []string
            gases           []string
            sounds          []string
            time            string
            temperature     string
            presence        string
            gas             string
            sound           string
        )

        for _, serie := range res.Values {
            if serie[0] != nil {
                time = serie[0].(string)
            } else {
                time = ""
            }
            if serie[1] != nil {
                temperature = serie[1].(string)
            } else {
                temperature = ""
            }
            if serie[2] != nil {
                presence = serie[2].(string)
            } else {
                presence = ""
            }
            if serie[3] != nil {
                gas = serie[3].(string)
            } else {
                gas = ""
            }
            if serie[4] != nil {
                sound = serie[4].(string)
            } else {
                sound = ""
            }

            times = append(times, time)
            temperatures = append(temperatures, temperature)
            presences = append(presences, presence)
            gases = append(gases, gas)
            sounds = append(sounds, sound)
        }

        agentData := AgentData{
            times,
            temperatures,
            presences,
            gases,
            sounds,
        }

        a := Agent{
            agent.Name,
            agentData,
        }

        data = append(data, a)
    }

	json.NewEncoder(w).Encode(data)
}

