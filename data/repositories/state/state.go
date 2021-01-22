package state

import (
	"errors"
	"github.com/coda-it/goutils/logger"
	"github.com/smart-evolution/shapi/data/persistence"
	"github.com/smart-evolution/shapi/domain/models/agent"
	"github.com/smart-evolution/shapi/domain/models/agent/types"
	"github.com/smart-evolution/shapi/domain/models/linux"
	modelState "github.com/smart-evolution/shapi/domain/models/state"
	"github.com/smart-evolution/shapi/domain/models/type1"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

const (
	collectionName = "state"
)

// Repository -
type Repository struct {
	persistence persistence.IPersistance
	loaded      bool
	model       modelState.State
}

// New - creates new instance of State
func New(p persistence.IPersistance, agents []agent.IAgent) *Repository {
	return &Repository{
		p,
		false,
		modelState.State{},
	}
}

func (s *Repository) findOneState(query interface{}) (modelState.State, error) {
	c := s.persistence.GetCollection(collectionName)

	var oneState modelState.State
	err := c.Find(query).One(&oneState)

	if err != nil {
		msg := "object not found"
		logger.Log(msg)
		return modelState.State{}, errors.New(msg)
	}

	return oneState, nil
}

func (s *Repository) load() {
	if s.loaded == false {
		persistedState, err := s.findOneState(bson.M{})

		if err != nil {
			logger.Log("failed to load state")
			return
		}

		s.model = persistedState
		s.loaded = true
	}
}

func (s *Repository) persist() {
	c := s.persistence.GetCollection(collectionName)
	_, err := c.Upsert(bson.M{}, s.model)

	if err != nil {
		logger.Log("failed to persist state")
	}
}

// SetIsAlerts - setter for `isAlerts`
func (s *Repository) SetIsAlerts(i bool) {
	s.model.IsAlerts = i
	s.persist()
}

// IsAlerts - getter for `isAlerts`
func (s *Repository) IsAlerts() bool {
	s.load()
	return s.model.IsAlerts
}

// SetSendAlert - setter for `sendAlert`
func (s *Repository) SetSendAlert(i bool) {
	s.model.SendAlert = i
	s.persist()
}

// SendAlert - getter for `sendAlert`
func (s *Repository) SendAlert() bool {
	s.load()
	return s.model.SendAlert
}

// AddAgent - adds agent to the memory state
func (s *Repository) AddAgent(id string, name string, ip string, agentType string) {
	s.load()
	_, err := s.AgentByID(id)

	if err == nil {
		logger.Log("not adding agent '" + id + "', agent already exists")
		return
	}

	logger.Log("adding home agent '" + name + "' with URL '" + ip + "'")
	rawType := strings.Split(agentType, "-")[0]

	if rawType == types.Type1 {
		a := type1.New(id, name, ip, agentType)
		s.model.Agents = append(s.model.Agents, a)
	} else if rawType == types.Linux {
		a := linux.New(id, name, ip, agentType)
		s.model.Agents = append(s.model.Agents, a)
	} else {
		a := agent.New(id, name, ip, agentType)
		s.model.Agents = append(s.model.Agents, a)
	}

	s.persist()
}

// Agents - returns list of available agents
func (s *Repository) Agents() []agent.IAgent {
	s.load()
	return s.model.Agents
}

// AgentByID - find corresponding agent by ID
func (s *Repository) AgentByID(id string) (agent.IAgent, error) {
	s.load()

	for _, ia := range s.model.Agents {
		switch a := ia.(type) {
		case *type1.Type1:
			if a.ID == id {
				return a, nil
			}
		}
	}

	return &agent.Agent{}, errors.New("no matching agent")
}

// AgentByIP - find corresponding agent by ID
func (s *Repository) AgentByIP(ip string) (agent.IAgent, error) {
	s.load()

	for _, ia := range s.model.Agents {
		a, ok := ia.(*agent.Agent)

		if !ok {
			return &agent.Agent{}, errors.New("type assertion error")
		}

		if a.IP == ip {
			return a, nil
		}
	}

	return &agent.Agent{}, errors.New("no matching agent")
}

// RemoveAgent - remove agent by ID
func (s *Repository) RemoveAgent(id string) error {
	for i, ia := range s.model.Agents {
		switch a := ia.(type) {
		case *type1.Type1:
			if a.ID == id {
				s.model.Agents = append(s.model.Agents[:i], s.model.Agents[i+1:]...)
				s.persist()
				return nil
			}
		case *agent.Agent:
			if a.ID == id {
				s.model.Agents = append(s.model.Agents[:i], s.model.Agents[i+1:]...)
				s.persist()
				return nil
			}
		}
	}

	return errors.New("no corresponding agent found")
}

// Reset - clears
func (s *Repository) Reset() {
	s.model.Agents = nil
	s.model.SendAlert = false
	s.model.IsAlerts = false
	s.persist()
}
