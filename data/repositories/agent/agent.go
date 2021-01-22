package agent

import (
	"errors"
	"github.com/coda-it/gowebserver/utils/logger"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/smart-evolution/shapi/data/dataflux"
	agentModel "github.com/smart-evolution/shapi/domain/models/agent"
)

const (
	databaseName = "smarthome"
)

// Repository -  agents repository
type Repository struct {
	Dataflux dataflux.IDataFlux
}

// New - creates agents repository instance
func New(d dataflux.IDataFlux) *Repository {
	return &Repository{
		d,
	}
}

// FetchType1Data - fetches data from agent-type-1
func (r *Repository) FetchType1Data(agentID string, period string) (agentModel.Type1DataJSON, error) {
	if r.Dataflux.IsConnected() != true {
		return agentModel.Type1DataJSON{}, errors.New("cannot feed data , Influx seems to be down")
	}

	measurement := "/.*/"

	if agentID != "" {
		measurement = agentID
	}

	q := client.Query{
		Command:  "SELECT time, temperature, presence, gas, sound, agent FROM \"" + measurement + "\" ORDER BY time DESC LIMIT " + period,
		Database: databaseName,
	}

	if r.Dataflux.IsConnected() != true {
		return agentModel.Type1DataJSON{}, errors.New("cannot feed data , Influx seems to be down")
	}

	resp, err := r.Dataflux.GetData(q)

	if err != nil {
		return agentModel.Type1DataJSON{}, err
	}

	if len(resp.Results) == 0 {
		return agentModel.Type1DataJSON{}, errors.New("not enough Results to fetch Series for measurement \"" + measurement + "\"")
	}

	series := resp.Results[0].Series

	var agentData agentModel.Type1DataJSON

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

		agentData = agentModel.Type1DataJSON{
			Time:        times,
			Temperature: temperatures,
			Presence:    presences,
			Gas:         gases,
			Sound:       sounds,
		}

		if err != nil {
			logger.Log(err.Error(), logger.ERROR)
		}
	}

	return agentData, nil
}
