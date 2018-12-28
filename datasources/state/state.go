package state

import (
    "errors"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/models/agent"
)

// IState - interface for datasource kept in the memory
type IState interface {
    SetIsAlerts(bool)
    IsAlerts() bool
    SetSendAlert(bool)
    SendAlert() bool
    Agents() []*agent.Agent
    AgentByID(string) (*agent.Agent, error)
}

// State - data source which keeps short data in memory
type State struct {
    isAlerts    bool
    sendAlert   bool
    agents      []*agent.Agent
}

// New - creates new instance of State
func New(agents []*agent.Agent) *State {
    return &State{
        isAlerts: false,
        sendAlert: false,
        agents: agents,
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
func (s *State) AddAgent(id string, name string, url string, agentType string) {
    utils.Log("adding home agent '" + name + "' with URL '" + url + "'")
    agent := agent.New(id, name, url, agentType)
    s.agents = append(s.agents, agent)
}

// Agents - returns list of available agents
func (s *State) Agents() []*agent.Agent {
    return s.agents
}

// AgentByID - find corresponding agent by ID
func (s *State) AgentByID(id string) (*agent.Agent, error) {
    for _, a := range s.agents {
        if a.ID() == id {
            return a, nil
        }
    }

    return  &agent.Agent{}, errors.New("No matching agent")
}
