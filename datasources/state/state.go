package state

import (
	"errors"
	"fmt"
	"github.com/smart-evolution/shapi/datasources/persistence"
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/models/agent/types"
	modelState "github.com/smart-evolution/shapi/models/state"
	"github.com/smart-evolution/shapi/models/type1"
	"github.com/smart-evolution/shapi/utils"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

// IState - interface for datasource kept in the memory
type IState interface {
	SetIsAlerts(bool)
	IsAlerts() bool
	SetSendAlert(bool)
	SendAlert() bool
	AddAgent(string, string, string, string)
	Agents() []agent.IAgent
	AgentByID(string) (agent.IAgent, error)
	AgentByIP(string) (agent.IAgent, error)
}

// State - data source which keeps short data in memory
type State struct {
	model  modelState.State
	src    *persistence.Persistance
	loaded bool
}

// New - creates new instance of State
func New(src *persistence.Persistance, agents []agent.IAgent) *State {
	model := modelState.State{
		IsAlerts:  false,
		SendAlert: false,
		Agents:    agents,
	}

	return &State{
		model,
		src,
		false,
	}
}

func (s *State) load() {
	c := s.src.GetCollection("state")

	var states []modelState.State
	err := c.Find(nil).All(&states)

	if err != nil {
		fmt.Println("----", err)
		utils.Log("failed to load `State`")
		return
	}

	if len(states) == 0 {
		utils.Log("state persistence empty")
	} else {
		s.model = states[0]
	}

	s.loaded = true
}

// SetIsAlerts - setted for `isAlerts`
func (s *State) SetIsAlerts(i bool) {
	s.model.IsAlerts = i

	c := s.src.GetCollection("state")
	_, err := c.Upsert(
		nil,
		s.model,
	)

	if err != nil {
		utils.Log("failed to persist `isAlerts`")
	}
}

// IsAlerts - getter for `isAlerts`
func (s *State) IsAlerts() bool {
	if s.loaded == false {
		s.load()
	}

	return s.model.IsAlerts
}

// SetSendAlert - setter for `sendAlert`
func (s *State) SetSendAlert(i bool) {
	s.model.SendAlert = i

	c := s.src.GetCollection("state")
	_, err := c.Upsert(
		bson.M{},
		s.model,
	)

	if err != nil {
		utils.Log("failed to persist `sendAlert`")
		return
	}
}

// SendAlert - getter for `sendAlert`
func (s *State) SendAlert() bool {
	if s.loaded == false {
		s.load()
	}

	return s.model.SendAlert
}

// AddAgent - adds agent to the memory state
func (s *State) AddAgent(id string, name string, ip string, agentType string) {
	utils.Log("adding home agent '" + name + "' with URL '" + ip + "'")
	rawType := strings.Split(agentType, "-")[0]

	if rawType == types.Type1 {
		a := type1.New(id, name, ip, agentType)
		s.model.Agents = append(s.model.Agents, a)
	} else {
		a := agent.New(id, name, ip, agentType)
		s.model.Agents = append(s.model.Agents, a)
	}

	c := s.src.GetCollection("state")
	//data, _ := bson.Marshal(s.model)
	//
	//var agents []agent.Agent
	//
	//for _, v := range s.model.Agents {
	//	agents = append(agents, *agent.New(v.ID, v.Name, v.IP, v.AgentType))
	//}
	//
	//data := struct{
	//	IsAlerts  bool			`bson:"isAlerts"`
	//	SendAlert bool			`bson:"sendAlert"`
	//	Agents    []agent.Agent `bson:"agents"`
	//}{
	//	IsAlerts: s.model.IsAlerts,
	//	SendAlert: s.model.SendAlert,
	//	Agents: agents,
	//}

	_, err := c.Upsert(
		bson.M{},
		s.model,
	)

	if err != nil {
		utils.Log("failed to persist `agent`")
	}
}

// Agents - returns list of available agents
func (s *State) Agents() []agent.IAgent {
	if s.loaded == false {
		s.load()
	}

	return s.model.Agents
}

// AgentByID - find corresponding agent by ID
func (s *State) AgentByID(id string) (agent.IAgent, error) {
	if s.loaded == false {
		s.load()
	}

	for _, ia := range s.model.Agents {
		a, ok := ia.(*agent.Agent)

		if !ok {
			return &agent.Agent{}, errors.New("type assertion error")
		}

		if a.ID == id {
			return ia, nil
		}
	}

	return &agent.Agent{}, errors.New("No matching agent")
}

// AgentByIP - find corresponding agent by ID
func (s *State) AgentByIP(ip string) (agent.IAgent, error) {
	if s.loaded == false {
		s.load()
	}

	for _, ia := range s.model.Agents {
		a, ok := ia.(*agent.Agent)

		if !ok {
			return &agent.Agent{}, errors.New("type assertion error")
		}

		if a.IP == ip {
			return a, nil
		}
	}

	return &agent.Agent{}, errors.New("No matching agent")
}
