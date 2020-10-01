package state

import (
	"github.com/coda-it/goutils/logger"
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/models/agent/types"
	"github.com/smart-evolution/shapi/models/type1"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

// AgentList - wrapper for slice of IAgent interface
type AgentList []agent.IAgent

// State - state model
type State struct {
	IsAlerts  bool `bson:"isAlerts"`
	SendAlert bool `bson:"sendAlert"`
	Agents    AgentList
}

// SetBSON - maps raw object to particular interface
func (al *AgentList) SetBSON(raw bson.Raw) error {
	var a []agent.Agent
	*al = AgentList{}

	err := raw.Unmarshal(&a)

	if err != nil {
		logger.Log("error setting BSON", err)
		return err
	}

	for _, v := range a {
		rawType := strings.Split(v.AgentType, "-")[0]

		switch rawType {
		case types.Type1:
			*al = append(*al, type1.New(v.ID, v.Name, v.IP, v.AgentType))
		default:
			*al = append(*al, agent.New(v.ID, v.Name, v.IP, v.AgentType))
		}
	}

	return nil
}
