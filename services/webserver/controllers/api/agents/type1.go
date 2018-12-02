package agents

import (
    "log"
    "errors"
    "github.com/smart-evolution/smarthome/state"
    "github.com/influxdata/influxdb/client/v2"
)

// Type1DataJSON - entity representing agent data
type Type1DataJSON struct {
    Time        []string  `json:"time"`
    Temperature []string  `json:"temperature"`
    Presence    []string  `json:"presence"`
    Gas         []string  `json:"gas"`
    Sound       []string  `json:"sound"`
}

// FetchType1 - fetches data for type1 agent
func FetchType1 (agentID string) ([]AgentJSON, error) {
    var type1Agents []AgentJSON

    if state.DataFlux.IsConnected() != true {
        return []AgentJSON{}, errors.New("cannot feed data , Influx seems to be down")
    }

    measurements := "/.*/"

    if agentID != "" {
        measurements = agentID
    }

    q := client.Query{
        Command: "SELECT time, temperature, presence, gas, sound, agent FROM " + measurements + " ORDER BY time DESC LIMIT 30",
        Database: "smarthome",
    }

    if state.DataFlux.IsConnected() != true {
        return []AgentJSON{}, errors.New("cannot feed data , Influx seems to be down")
    }

    resp, err := state.DataFlux.GetData(q)

    if err != nil {
        return []AgentJSON{}, err
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

        agentData := Type1DataJSON{
            times,
            temperatures,
            presences,
            gases,
            sounds,
        }

        if err != nil {
            log.Println("services: ", err)
        }

        a := AgentJSON{
            agentID,
            agentName,
            agentData,
            "type1",
        }

        type1Agents = append(type1Agents, a)
    }

    return type1Agents, nil
}
