package agent

import "strings"

// IAgent - interface for Agent
type IAgent interface {
	ID() string
	SetID(string)
	Name() string
	SetName(string)
	IP() string
	SetIP(string)
	AgentType() string
	RawType() string
	SetAgentType(string)
	IsOnline() bool
	SetIsOnline(bool)
}

// Agent - hardware entity
type Agent struct {
	iD        string
	name      string
	iP        string
	agentType string
	isOnline  bool
}

// New - creates new entity of Agent
func New(id string, name string, ip string, agentType string) *Agent {
	return &Agent{
		iD:        id,
		name:      name,
		iP:        ip,
		agentType: agentType,
		isOnline:  true,
	}
}

// ID - iD getter
func (a *Agent) ID() string {
	return a.iD
}

// SetID - iD setter
func (a *Agent) SetID(id string) {
	a.iD = id
}

// Name - name getter
func (a *Agent) Name() string {
	return a.name
}

// SetName - name setter
func (a *Agent) SetName(name string) {
	a.name = name
}

// IP - iP getter
func (a *Agent) IP() string {
	return a.iP
}

// SetIP - iP setter
func (a *Agent) SetIP(id string) {
	a.iP = id
}

// AgentType - agentType getter
func (a *Agent) AgentType() string {
	return a.agentType
}

// SetAgentType - agentType setter
func (a *Agent) SetAgentType(agentType string) {
	a.agentType = agentType
}

// RawType - gets unversioned agentType getter
func (a *Agent) RawType() string {
	return strings.Split(a.agentType, "-")[0]
}

// IsOnline - isOnline getter
func (a *Agent) IsOnline() bool {
	return a.isOnline
}

// SetIsOnline - isOnline setter
func (a *Agent) SetIsOnline(isOnline bool) {
	a.isOnline = isOnline
}
