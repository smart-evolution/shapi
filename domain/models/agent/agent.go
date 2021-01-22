package agent

import "strings"

// IAgent - interface for Agent
type IAgent interface {
	RawType() string
}

// Agent - hardware entity
type Agent struct {
	ID        string `bson:"_id"`
	Name      string `bson:"name"`
	IP        string `bson:"ip"`
	AgentType string `bson:"agentType"`
	IsOnline  bool   `bson:"isOnline"`
}

// New - creates new entity of Agent
func New(id string, name string, ip string, agentType string) *Agent {
	return &Agent{
		ID:        id,
		Name:      name,
		IP:        ip,
		AgentType: agentType,
		IsOnline:  true,
	}
}

// RawType - gets unversioned agentType getter
func (a *Agent) RawType() string {
	return strings.Split(a.AgentType, "-")[0]
}
