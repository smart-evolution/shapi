package main

import (
	"os"
    "log"
    "sync"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/state"
    "github.com/smart-evolution/smarthome/services/homebot"
    "github.com/smart-evolution/smarthome/services/dataflux"
    "github.com/smart-evolution/smarthome/services/email"
    "github.com/smart-evolution/smarthome/services/persistence"
    "github.com/smart-evolution/smarthome/services/webserver"
    "github.com/smart-evolution/smarthome/models/user"
    "gopkg.in/mgo.v2/bson"
)

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./version.go

func main() {
    var wg sync.WaitGroup
    utils.VERSION = VERSION

    p := persistence.New(os.Getenv("MONGOLAB_URI"), os.Getenv("DB_NAME"))
    state.Persistance = p

    df := dataflux.New("http://localhost:8086")
    state.DataFlux = df

    m := email.New(os.Getenv("EMAILNAME"), os.Getenv("EMAILPASS"), os.Getenv("SMTPPORT"), os.Getenv("SMTPAUTHURL"))
    state.Mailer = m
    c := state.Persistance.GetCollection("users")

    var users []user.User
    err := c.Find(bson.M{}).All(&users)

    if err != nil {
        log.Println("services: Alert recipients not found", err)
    }

    for _, u := range users {
        state.Mailer.AddRecipient(u.Username)
    }

    hb := homebot.New("hardware/agents.config", df, m)
    state.HomeBot = hb
    go hb.RunService(wg)

    ws := webserver.New(os.Getenv("PORT"), df)
    go ws.RunService(wg)

    wg.Wait()
}

