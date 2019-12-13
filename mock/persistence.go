package mock

import (
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/models/agent/types"
	"github.com/smart-evolution/shapi/models/state"
	"github.com/smart-evolution/shapi/models/user"
	"gopkg.in/mgo.v2"
)

// PersistanceMock - data source keeping system state and user data mock
type PersistanceMock struct {
	session *mgo.Session
	dbName  string
}

// NewPersistanceMock - creates new persistence mock
func NewPersistanceMock(dbURI string, dbName string) *PersistanceMock {
	return &PersistanceMock{
		&mgo.Session{},
		dbName,
	}
}

// FindOneUser - find one user entry
func (p *PersistanceMock) FindOneUser(query interface{}) (user.User, error) {
	var u user.User
	return u, nil
}

// FindAllUsers - find all users
func (p *PersistanceMock) FindAllUsers(query interface{}) ([]user.User, error) {
	var users []user.User
	return users, nil
}

// FindOneAgentConfig - find one agentConfig entry
func (p *PersistanceMock) FindOneAgentConfig(query interface{}) (agent.Config, error) {
	var cnf agent.Config
	return cnf, nil
}

// FindAllAgentConfigs - find all agentConfigs
func (p *PersistanceMock) FindAllAgentConfigs(query interface{}) ([]agent.Config, error) {
	var configs []agent.Config
	return configs, nil
}

// FindOneState - find one state entry
func (p *PersistanceMock) FindOneState(query interface{}) (state.State, error) {
	agent1 := agent.New("livingroom", "Living room", "192.168.1.2", types.Type1)
	agent2 := agent.New("bedroom", "Bed room", "192.168.1.3", types.Type2)
	agents := []agent.IAgent{agent1, agent2}

	s := state.State{
		IsAlerts:  false,
		SendAlert: false,
		Agents:    agents,
	}
	return s, nil
}

// Insert - instert any object into persistence
func (p *PersistanceMock) Insert(collection string, docs ...interface{}) error {
	return nil
}

// Upsert - insert and if exists update entry in persistence
func (p *PersistanceMock) Upsert(collection string, selector interface{}, update interface{}) error {
	return nil
}
