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

// CtrHome - controller for retrieving agents list data
func CtrAgents(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    var data []Agent

    if services.InfluxConnected != true {
        w.WriteHeader(http.StatusInternalServerError)
        log.Println("services: cannot feed data , Influx seems to be down")
        return
    }

    q := client.Query{
        Command:    "SELECT time, temperature, presence, gas, sound, agent FROM /.*/ ORDER BY time DESC LIMIT 30",
        Database:   "smarthome",
    }

    resp, err := services.InfluxClient.Query(q)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        log.Println("services: ", err)
        return
    }

    if len(resp.Results) == 0 {
        w.WriteHeader(http.StatusNoContent)
        return
    }

    series := resp.Results[0].Series

    for _, agent := range series {
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
            agentName       string
        )

        agentID := agent.Name

        for _, serie := range agent.Values {
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
            agentName = serie[5].(string)
        }

        agentData := AgentData{
            times,
            temperatures,
            presences,
            gases,
            sounds,
        }

        a := Agent{
            agentID,
            agentName,
            agentData,
        }

        data = append(data, a)
    }

    json.NewEncoder(w).Encode(data)
}

