package persistence

import (
	"errors"
	"github.com/coda-it/goutils/logger"
	"gopkg.in/mgo.v2"
)

// IPersistance - interface for user settings and general purpose storage
type IPersistance interface {
	GetCollection(string) *mgo.Collection
	DropDatabase() error
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

func (p *Persistance) getDatabase() *mgo.Database {
	return p.session.DB(p.dbName)
}

// GetCollection - gets collection from Persistance instance
func (p *Persistance) GetCollection(name string) *mgo.Collection {
	ds := p.getDatabase()
	return ds.C(name)
}

// DropDatabase - clear whole database
func (p *Persistance) DropDatabase() error {
	ds := p.getDatabase()
	return ds.DropDatabase()
}

// Insert - instert any object into persistence
func (p *Persistance) Insert(collection string, docs ...interface{}) error {
	c := p.GetCollection(collection)

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
	c := p.GetCollection(collection)

	_, err := c.Upsert(selector, update)

	if err != nil {
		msg := "object not upserted"
		logger.Log(msg, err)
		return errors.New(msg)
	}

	return nil
}
