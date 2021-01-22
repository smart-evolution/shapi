package agent

import (
	agentModel "github.com/smart-evolution/shapi/domain/models/agent"
)

// IStateRepository -
type IStateRepository interface {
	AddAgent(id string, name string, ip string, agentType string)
	AgentByID(id string) (agentModel.IAgent, error)
	RemoveAgent(id string) error
	Agents() []agentModel.IAgent
	AgentByIP(ip string) (agentModel.IAgent, error)
}

// IAgentRepository -
type IAgentRepository interface {
	FetchType1Data(agentID string, period string) (agentModel.Type1DataJSON, error)
}

// IAgentConfigsRepository -
type IAgentConfigsRepository interface {
	FindOneAgentConfig(query interface{}) (agentModel.Config, error)
	FindAllAgentConfigs(query interface{}) ([]agentModel.Config, error)
	UpdateAgentConfigs(agentID string, config interface{}) error
}
