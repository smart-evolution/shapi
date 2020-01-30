package linux

import (
	"github.com/smart-evolution/shapi/models/agent"
)

// ILinux - interface for ILinux hardware
type ILinux interface {
	agent.IAgent
}

// Linux - hardware entity
type Linux struct {
	agent.Agent `bson:",inline"`
}

// New - creates new entity of Type1 Agent
func New(id string, name string, ip string, agentType string) *Linux {
	a := agent.New(id, name, ip, agentType)

	return &Linux{
		Agent: *a,
	}
}
