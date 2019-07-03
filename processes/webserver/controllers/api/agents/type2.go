package agents

import (
	"github.com/smart-evolution/smarthome/models/agent"
	"github.com/smart-evolution/smarthome/models/agent/types"
)

// Type2DataJSON - entity representing agent data
type Type2DataJSON struct {
}

// FetchType2 - fetches data for type2 agent
func FetchType2(agentID string, agents []agent.IAgent) ([]AgentJSON, error) {
	var agentsJSON []AgentJSON

	for _, a := range agents {
		if a.AgentType() == types.TYPE2 {
			agent := AgentJSON{
				ID:        a.ID(),
				Name:      a.Name(),
				Data:      struct{}{},
				AgentType: types.TYPE2,
				IP:        a.IP(),
			}

			agentsJSON = append(agentsJSON, agent)
		}
	}

	return agentsJSON, nil
}
