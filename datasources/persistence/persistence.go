package persistence

import (
    "log"
    "gopkg.in/mgo.v2"
)

// Persistance - data source keeping system state and user data
type Persistance struct {
    session *mgo.Session
    dbName  string
}

// New - creates new instance of Persistance
func New(dbURI string, dbName string) Persistance {
    log.Println("Connecting to mgo with URI = " + dbURI)
    session, err := mgo.Dial(dbURI)
    session.SetMode(mgo.Monotonic, true)

    if err != nil {
        panic(err)
    }

    return Persistance{
        session,
        dbName,
    }
}

func (p Persistance) getDatabase() *mgo.Database {
    return p.session.DB(p.dbName)
}

// GetCollection - gets collection from Persistance instance
func (p Persistance) GetCollection(name string) *mgo.Collection {
    ds := p.getDatabase()
    return ds.C(name)
}
