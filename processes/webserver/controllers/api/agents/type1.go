package agents

import (
	"errors"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/smart-evolution/smarthome/datasources/dataflux"
	"github.com/smart-evolution/smarthome/utils"
)

// Type1DataJSON - entity representing agent data
type Type1DataJSON struct {
	Time        []string `json:"time"`
	Temperature []string `json:"temperature"`
	Presence    []string `json:"presence"`
	Gas         []string `json:"gas"`
	Sound       []string `json:"sound"`
}

// FetchType1Data - fetches data for type1 agent
func FetchType1Data(agentID string, period string, df dataflux.IDataFlux) (Type1DataJSON, error) {
	if df.IsConnected() != true {
		return Type1DataJSON{}, errors.New("cannot feed data , Influx seems to be down")
	}

	measurement := "/.*/"

	if agentID != "" {
		measurement = agentID
	}

	q := client.Query{
		Command:  "SELECT time, temperature, presence, gas, sound, agent FROM \"" + measurement + "\" ORDER BY time DESC LIMIT " + period,
		Database: "smarthome",
	}

	if df.IsConnected() != true {
		return Type1DataJSON{}, errors.New("cannot feed data , Influx seems to be down")
	}

	resp, err := df.GetData(q)

	if err != nil {
		return Type1DataJSON{}, err
	}

	if len(resp.Results) == 0 {
		return Type1DataJSON{}, errors.New("not enough Results to fetch Series for measurement \"" + measurement + "\"")
	}

	series := resp.Results[0].Series

	var agentData Type1DataJSON

	for _, agent := range series {
		var (
			times        []string
			temperatures []string
			presences    []string
			gases        []string
			sounds       []string
			time         string
			temperature  string
			presence     string
			gas          string
			sound        string
		)

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
		}

		agentData = Type1DataJSON{
			times,
			temperatures,
			presences,
			gases,
			sounds,
		}

		if err != nil {
			utils.Log(err)
		}
	}

	return agentData, nil
}
