package api

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/smart-evolution/smarthome/services"
    "github.com/smart-evolution/smarthome/models"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
    "github.com/influxdata/influxdb/client/v2"
)

// Agent - entity representing agent state
type Agent struct {
    ID          string      `json:"id"`
    Name        string      `json:"name"`
    Data        AgentData   `json:"data"`
    AgentType   string      `json:"type"`
}

// AgentData - entity representing agent data
type AgentData struct {
    Time        []string  `json:"time"`
    Temperature []string  `json:"temperature"`
    Presence    []string  `json:"presence"`
    Gas         []string  `json:"gas"`
    Sound       []string  `json:"sound"`
}

// CtrHome - controller for retrieving agent data
func CtrHome(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    var data []Agent

    agentID := opt.Params["agent"]

    if services.InfluxConnected != true {
        w.WriteHeader(http.StatusInternalServerError)
        log.Println("services: cannot feed data , Influx seems to be down")
        return
    }

    q := client.Query{
        Command:    "SELECT time, temperature, presence, gas, sound, agent FROM " + agentID + " ORDER BY time DESC LIMIT 30",
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
        agentName       string
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
        agentName = serie[5].(string)
    }

    agentData := AgentData{
        times,
        temperatures,
        presences,
        gases,
        sounds,
    }

    agentInRegistry, err := models.FindAgentById(agentID)

    if err != nil {
        log.Println("services: ", err)
    }

    a := Agent{
        agentID,
        agentName,
        agentData,
        agentInRegistry.AgentType,
    }

    data = append(data, a)

    json.NewEncoder(w).Encode(data)
}

