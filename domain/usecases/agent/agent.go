package agent

import (
	agentModel "github.com/smart-evolution/shapi/domain/models/agent"
)

// Usecase - agent usecases
type Usecase struct {
	stateRepository        IStateRepository
	agentRepository        IAgentRepository
	agentConfigsRepository IAgentConfigsRepository
}

// New - creates new agent usecases
func New(sr IStateRepository, ar IAgentRepository, acr IAgentConfigsRepository) *Usecase {
	return &Usecase{
		sr,
		ar,
		acr,
	}
}

// AddAgent - adds agent to state
func (u *Usecase) AddAgent(id string, name string, ip string, agentType string) {
	u.stateRepository.AddAgent(id, name, ip, agentType)
}

// RemoveAgent - removes agent from repository
func (u *Usecase) RemoveAgent(id string) error {
	return u.stateRepository.RemoveAgent(id)
}

// AgentByID - gets agent by id
func (u *Usecase) AgentByID(id string) (agentModel.IAgent, error) {
	return u.stateRepository.AgentByID(id)
}

// Agents - gets all agents
func (u *Usecase) Agents() []agentModel.IAgent {
	return u.stateRepository.Agents()
}

// FetchType1Data - fetches data for agent type 1
func (u *Usecase) FetchType1Data(agentID string, period string) (agentModel.Type1DataJSON, error) {
	return u.agentRepository.FetchType1Data(agentID, period)
}

// FindOneAgentConfig - gets agent config
func (u *Usecase) FindOneAgentConfig(query interface{}) (agentModel.Config, error) {
	return u.agentConfigsRepository.FindOneAgentConfig(query)
}

// FindAllAgentConfigs - gets all agent configs
func (u *Usecase) FindAllAgentConfigs(query interface{}) ([]agentModel.Config, error) {
	return u.agentConfigsRepository.FindAllAgentConfigs(query)
}

// UpdateAgentConfigs - updates agent config
func (u *Usecase) UpdateAgentConfigs(agentID string, config interface{}) error {
	return u.agentConfigsRepository.UpdateAgentConfigs(agentID, config)
}
