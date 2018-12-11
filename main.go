package main

import (
	"os"
    "log"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/datasources/persistence"
    "github.com/smart-evolution/smarthome/datasources/dataflux"
    "github.com/smart-evolution/smarthome/datasources/state"
    "github.com/smart-evolution/smarthome/processes/homebot"
    "github.com/smart-evolution/smarthome/services/email"
    "github.com/smart-evolution/smarthome/processes/webserver"
    "github.com/smart-evolution/smarthome/models/user"
    "gopkg.in/mgo.v2/bson"
)

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./version.go

func main() {
    utils.VERSION = VERSION

    s := state.New()
    s.SetupAgents("hardware/agents.config")

    p := persistence.New(os.Getenv("MONGOLAB_URI"), os.Getenv("DB_NAME"))
    df := dataflux.New("http://localhost:8086")
    m := email.New(os.Getenv("EMAILNAME"), os.Getenv("EMAILPASS"), os.Getenv("SMTPPORT"), os.Getenv("SMTPAUTHURL"))

    c := p.GetCollection("users")
    var users []user.User
    err := c.Find(bson.M{}).All(&users)

    if err != nil {
        log.Println("services: Alert recipients not found", err)
    }

    for _, u := range users {
        m.AddRecipient(u.Username)
    }

    hb := homebot.New(df, m, s)
    go hb.RunService()

    ws := webserver.New(os.Getenv("PORT"), df, p, s)
    ws.RunService()
}

