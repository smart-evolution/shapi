package agent

// Agent - entity representing agent state
type Agent struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Data      interface{} `json:"data"`
	AgentType string      `json:"type"`
	IP        string      `json:"ip"`
	IsOnline  bool        `json:"isOnline"`
}
