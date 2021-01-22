package agentconfigs

import (
	"errors"
	"github.com/coda-it/gowebserver/utils/logger"
	"github.com/smart-evolution/shapi/data/persistence"
	"github.com/smart-evolution/shapi/domain/models/agent"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName = "agentConfigs"
)

// Repository - agent configs repository
type Repository struct {
	persistence persistence.IPersistance
}

// New - creates new agent configs repository
func New(p persistence.IPersistance) *Repository {
	return &Repository{
		p,
	}
}

// FindOneAgentConfig - find one agentConfig entry
func (r *Repository) FindOneAgentConfig(query interface{}) (agent.Config, error) {
	c := r.persistence.GetCollection("users")

	var cnf agent.Config
	err := c.Find(query).One(&cnf)

	if err != nil {
		msg := "object not found"
		logger.Log(msg, logger.ERROR)
		return agent.Config{}, errors.New(msg)
	}

	return cnf, nil
}

// FindAllAgentConfigs - find all agentConfigs
func (r *Repository) FindAllAgentConfigs(query interface{}) ([]agent.Config, error) {
	c := r.persistence.GetCollection(collectionName)

	var configs []agent.Config
	err := c.Find(query).All(&configs)

	if err != nil {
		msg := "objects not found"
		logger.Log(msg, logger.ERROR)
		return nil, errors.New(msg)
	}

	return configs, nil
}

// UpdateAgentConfigs - updates agent config
func (r *Repository) UpdateAgentConfigs(agentID string, config interface{}) error {
	c := r.persistence.GetCollection(collectionName)
	_, err := c.Upsert(bson.M{"agentId": agentID}, config)
	return err
}
