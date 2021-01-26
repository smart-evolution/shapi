package addedagent

// AddedAgent - added agent JSON
type AddedAgent struct {
	ID   string `json:"agentID"`
	Name string `json:"agentName"`
	IP   string `json:"agentIP"`
	Type string `json:"agentType"`
}
