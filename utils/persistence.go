package utils

import (
    "os"
    "gopkg.in/mgo.v2"
)

var dBSession *mgo.Session

func SetSession(session *mgo.Session) {
    dBSession = session
}

func GetDataSource() *mgo.Database {
    dbName := os.Getenv("DB_NAME")
    return dBSession.DB(dbName)
}
