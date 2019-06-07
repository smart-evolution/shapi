package state

import (
	"errors"
	"github.com/smart-evolution/smarthome/models/agent"
	"github.com/smart-evolution/smarthome/utils"
)

// IState - interface for datasource kept in the memory
type IState interface {
	SetIsAlerts(bool)
	IsAlerts() bool
	SetSendAlert(bool)
	SendAlert() bool
	Agents() []agent.IAgent
	AgentByID(string) (agent.IAgent, error)
}

// State - data source which keeps short data in memory
type State struct {
	isAlerts  bool
	sendAlert bool
	agents    []agent.IAgent
}

// New - creates new instance of State
func New(agents []agent.IAgent) *State {
	return &State{
		isAlerts:  false,
		sendAlert: false,
		agents:    agents,
	}
}

// SetIsAlerts - setted for `isAlerts`
func (s *State) SetIsAlerts(i bool) {
	s.isAlerts = i
}

// IsAlerts - getter for `isAlerts`
func (s *State) IsAlerts() bool {
	return s.isAlerts
}

// SetSendAlert - setter for `sendAlert`
func (s *State) SetSendAlert(i bool) {
	s.sendAlert = i
}

// SendAlert - getter for `sendAlert`
func (s *State) SendAlert() bool {
	return s.sendAlert
}

// AddAgent - adds agent to the memory state
func (s *State) AddAgent(id string, name string, ip string, agentType string) {
	utils.Log("adding home agent '" + name + "' with URL '" + ip + "'")
	agent := agent.New(id, name, ip, agentType)
	s.agents = append(s.agents, agent)
}

// Agents - returns list of available agents
func (s *State) Agents() []agent.IAgent {
	return s.agents
}

// AgentByID - find corresponding agent by ID
func (s *State) AgentByID(id string) (agent.IAgent, error) {
	for _, a := range s.agents {
		if a.ID() == id {
			return a, nil
		}
	}

	return &agent.Agent{}, errors.New("No matching agent")
}
