package agents

import (
    "github.com/smart-evolution/smarthome/models"
)

// Type2DataJSON - entity representing agent data
type Type2DataJSON struct {
}

// FetchType2 - fetches data for type2 agent
func FetchType2 (agentID string) ([]AgentJSON, error) {
    var agents []AgentJSON

    for _, a := range models.Agents {
        if a.AgentType == "type2" {
            agent := AgentJSON{
                ID: a.ID,
                Name: a.Name,
                Data: struct{}{},
                AgentType: "type2",
            }

            agents = append(agents, agent)
        }
    }

    return agents, nil
}
