package agents

// AgentJSON - entity representing agent state
type AgentJSON struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Data      interface{} `json:"data"`
	AgentType string      `json:"type"`
	IP		  string	  `json:"ip"`
}
