package agents

import (
    "github.com/smart-evolution/smarthome/models/agent"
)

// Type2DataJSON - entity representing agent data
type Type2DataJSON struct {
}

// FetchType2 - fetches data for type2 agent
func FetchType2 (agentID string, agents []*agent.Agent) ([]AgentJSON, error) {
    var agentsJSON []AgentJSON

    for _, a := range agents {
        if a.AgentType() == "type2" {
            agent := AgentJSON{
                ID: a.ID(),
                Name: a.Name(),
                Data: struct{}{},
                AgentType: "type2",
            }

            agentsJSON = append(agentsJSON, agent)
        }
    }

    return agentsJSON, nil
}
