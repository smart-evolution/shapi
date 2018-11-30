package main

import (
	"os"
	"fmt"
    "log"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/services/homebot"
    "github.com/smart-evolution/smarthome/services/dataflux"
    "github.com/smart-evolution/smarthome/services/email"
    "github.com/smart-evolution/smarthome/services/persistence"
    "github.com/smart-evolution/smarthome/models/user"
	"github.com/smart-evolution/smarthome/controllers"
	"github.com/smart-evolution/smarthome/controllers/api"
	"github.com/coda-it/gowebserver"
    "gopkg.in/mgo.v2/bson"
)

func getServerAddress() (string, error) {
	port := os.Getenv("PORT")

	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./version.go

func main() {
    addr, _ := getServerAddress()
    utils.VERSION = VERSION

    serverOptions := gowebserver.WebServerOptions{
        Port: addr,
        StaticFilesUrl: "/static/",
        StaticFilesDir: "public",
    }

    utils.Persistance = persistence.New(os.Getenv("MONGOLAB_URI"), os.Getenv("DB_NAME"))
    utils.DataFlux = dataflux.New("http://localhost:8086")
    hb := homebot.New("hardware/agents.config")
    go hb.RunHomeService()

    utils.Mailer = email.New(os.Getenv("EMAILNAME"), os.Getenv("EMAILPASS"), os.Getenv("SMTPPORT"), os.Getenv("SMTPAUTHURL"))
    ds := utils.Persistance.GetDatabase()
    c := ds.C("users")

    var users []user.User
    err := c.Find(bson.M{}).All(&users)

    if err != nil {
        log.Println("services: Alert recipients not found", err)
    }

    for _, u := range users {
        utils.Mailer.AddRecipient(u.Username)
    }

    server := gowebserver.New(serverOptions, controllers.NotFound)
    server.Router.AddRoute("/login/register", controllers.Register)
    server.Router.AddRoute("/login/logout", controllers.AuthenticateLogout)
    server.Router.AddRoute("/login", controllers.Authenticate)
    server.Router.AddRoute("/", controllers.CtrDashboard)
    server.Router.AddRoute("/api/agents", api.CtrAgents)
    server.Router.AddRoute("/api/agents/{agent}", api.CtrAgents)
    server.Router.AddRoute("/api/alerts", api.CtrAlerts)
    server.Router.AddRoute("/api/sendalert", api.CtrSendAlert)

    server.Run()
}

