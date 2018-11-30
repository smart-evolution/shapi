package persistence

import (
    "log"
    "gopkg.in/mgo.v2"
)

type Persistance struct {
    session *mgo.Session
    dbName  string
}

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

func (p *Persistance) GetDatabase() *mgo.Database {
    return p.session.DB(p.dbName)
}

