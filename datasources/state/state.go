package state

import (
	"errors"
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
	RemoveAgent(string) error
}

// State - data source which keeps short data in memory
type State struct {
	model       modelState.State
	persistence persistence.IPersistance
	loaded      bool
}

// New - creates new instance of State
func New(p persistence.IPersistance, agents []agent.IAgent) *State {
	model := modelState.State{
		IsAlerts:  false,
		SendAlert: false,
		Agents:    agents,
	}

	return &State{
		model,
		p,
		false,
	}
}

func (s *State) load() {
	if s.loaded == false {
		persistedState, err := s.persistence.FindOneState(bson.M{})

		if err != nil {
			utils.Log("failed to load state")
			return
		}

		s.model = persistedState
		s.loaded = true
	}
}

func (s *State) persist() {
	err := s.persistence.Upsert("state", bson.M{}, s.model)

	if err != nil {
		utils.Log("failed to persist state")
	}
}

// SetIsAlerts - setter for `isAlerts`
func (s *State) SetIsAlerts(i bool) {
	s.model.IsAlerts = i
	s.persist()
}

// IsAlerts - getter for `isAlerts`
func (s *State) IsAlerts() bool {
	s.load()
	return s.model.IsAlerts
}

// SetSendAlert - setter for `sendAlert`
func (s *State) SetSendAlert(i bool) {
	s.model.SendAlert = i
	s.persist()
}

// SendAlert - getter for `sendAlert`
func (s *State) SendAlert() bool {
	s.load()
	return s.model.SendAlert
}

// AddAgent - adds agent to the memory state
func (s *State) AddAgent(id string, name string, ip string, agentType string) {
	s.load()
	_, err := s.AgentByID(id)

	if err == nil {
		utils.Log("not adding agent '" + id + "', agent already exists")
		return
	}

	utils.Log("adding home agent '" + name + "' with URL '" + ip + "'")
	rawType := strings.Split(agentType, "-")[0]

	if rawType == types.Type1 {
		a := type1.New(id, name, ip, agentType)
		s.model.Agents = append(s.model.Agents, a)
	} else {
		a := agent.New(id, name, ip, agentType)
		s.model.Agents = append(s.model.Agents, a)
	}

	s.persist()
}

// Agents - returns list of available agents
func (s *State) Agents() []agent.IAgent {
	s.load()
	return s.model.Agents
}

// AgentByID - find corresponding agent by ID
func (s *State) AgentByID(id string) (agent.IAgent, error) {
	s.load()

	for _, ia := range s.model.Agents {
		a, ok := ia.(*agent.Agent)

		if !ok {
			return &agent.Agent{}, errors.New("type assertion error")
		}

		if a.ID == id {
			return ia, nil
		}
	}

	return &agent.Agent{}, errors.New("no matching agent")
}

// AgentByIP - find corresponding agent by ID
func (s *State) AgentByIP(ip string) (agent.IAgent, error) {
	s.load()

	for _, ia := range s.model.Agents {
		a, ok := ia.(*agent.Agent)

		if !ok {
			return &agent.Agent{}, errors.New("type assertion error")
		}

		if a.IP == ip {
			return a, nil
		}
	}

	return &agent.Agent{}, errors.New("no matching agent")
}

// RemoveAgent - remove agent by ID
func (s *State) RemoveAgent(id string) error {
	for i, ia := range s.model.Agents {
		switch a := ia.(type) {
		case *type1.Type1:
			if a.ID == id {
				s.model.Agents = append(s.model.Agents[:i], s.model.Agents[i+1:]...)
				s.persist()
				return nil
			}
		}

	}

	return errors.New("no corresponding agent found")
}
