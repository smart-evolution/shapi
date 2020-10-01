package persistence

import (
	"errors"
	"github.com/coda-it/goutils/logger"
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/models/state"
	"github.com/smart-evolution/shapi/models/user"
	"gopkg.in/mgo.v2"
)

// IPersistance - interface for user settings and general purpose storage
type IPersistance interface {
	FindOneUser(interface{}) (user.User, error)
	FindAllUsers(interface{}) ([]user.User, error)
	FindOneAgentConfig(interface{}) (agent.Config, error)
	FindAllAgentConfigs(interface{}) ([]agent.Config, error)
	FindOneState(interface{}) (state.State, error)
	Insert(string, ...interface{}) error
	Upsert(string, interface{}, interface{}) error
}

// Persistance - data source keeping system state and user data
type Persistance struct {
	session *mgo.Session
	dbName  string
}

// New - creates new instance of Persistance
func New(dbURI string, dbName string) *Persistance {
	logger.Log("connecting to mgo with URI '" + dbURI + "'")
	session, err := mgo.Dial(dbURI)
	session.SetMode(mgo.Monotonic, true)

	if err != nil {
		logger.Log("error dialing mongodb", err)
		panic(err)
	}

	return &Persistance{
		session,
		dbName,
	}
}

func (p *Persistance) getCollection(name string) *mgo.Collection {
	return p.session.DB(p.dbName).C(name)
}

// FindOneUser - find one user entry
func (p *Persistance) FindOneUser(query interface{}) (user.User, error) {
	c := p.getCollection("users")

	var u user.User
	err := c.Find(query).One(&u)

	if err != nil {
		msg := "object not found"
		logger.Log(msg)
		return user.User{}, errors.New(msg)
	}

	return u, nil
}

// FindAllUsers - find all users
func (p *Persistance) FindAllUsers(query interface{}) ([]user.User, error) {
	c := p.getCollection("users")

	var users []user.User
	err := c.Find(query).All(&users)

	if err != nil {
		msg := "objects not found"
		logger.Log(msg)
		return nil, errors.New(msg)
	}

	return users, nil
}

// FindOneAgentConfig - find one agentConfig entry
func (p *Persistance) FindOneAgentConfig(query interface{}) (agent.Config, error) {
	c := p.getCollection("users")

	var cnf agent.Config
	err := c.Find(query).One(&cnf)

	if err != nil {
		msg := "object not found"
		logger.Log(msg)
		return agent.Config{}, errors.New(msg)
	}

	return cnf, nil
}

// FindAllAgentConfigs - find all agentConfigs
func (p *Persistance) FindAllAgentConfigs(query interface{}) ([]agent.Config, error) {
	c := p.getCollection("agentConfigs")

	var configs []agent.Config
	err := c.Find(query).All(&configs)

	if err != nil {
		msg := "objects not found"
		logger.Log(msg)
		return nil, errors.New(msg)
	}

	return configs, nil
}

// FindOneState - find one state entry
func (p *Persistance) FindOneState(query interface{}) (state.State, error) {
	c := p.getCollection("state")

	var s state.State
	err := c.Find(query).One(&s)

	if err != nil {
		msg := "object not found"
		logger.Log(msg)
		return state.State{}, errors.New(msg)
	}

	return s, nil
}

// Insert - instert any object into persistence
func (p *Persistance) Insert(collection string, docs ...interface{}) error {
	c := p.getCollection(collection)

	err := c.Insert(docs...)

	if err != nil {
		msg := "object not inserted"
		logger.Log(msg, err)
		return errors.New(msg)
	}

	return nil
}

// Upsert - insert and if exists update entry in persistence
func (p *Persistance) Upsert(collection string, selector interface{}, update interface{}) error {
	c := p.getCollection(collection)

	_, err := c.Upsert(selector, update)

	if err != nil {
		msg := "object not upserted"
		logger.Log(msg, err)
		return errors.New(msg)
	}

	return nil
}
