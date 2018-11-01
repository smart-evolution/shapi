package utils

import (
    "os"
    "gopkg.in/mgo.v2"
)

var dBSession *mgo.Session

// SetSession - set mongodb session
func SetSession(session *mgo.Session) {
    dBSession = session
}

// GetDataSource - get mongodb data source
func GetDataSource() *mgo.Database {
    dbName := os.Getenv("DB_NAME")
    return dBSession.DB(dbName)
}
