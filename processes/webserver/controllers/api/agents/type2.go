package agents

import (
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/models/agent/types"
)

// Type2DataJSON - entity representing agent data
type Type2DataJSON struct {
}

// FetchType2 - fetches data for type2 agent
func FetchType2(agentID string, agents []agent.IAgent) ([]AgentJSON, error) {
	var agentsJSON []AgentJSON

	for _, a := range agents {
		if a.AgentType() == types.Type2 {
			agent := AgentJSON{
				ID:        a.ID(),
				Name:      a.Name(),
				Data:      struct{}{},
				AgentType: types.Type2,
				IP:        a.IP(),
			}

			agentsJSON = append(agentsJSON, agent)
		}
	}

	return agentsJSON, nil
}
