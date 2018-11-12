package main

import (
	"os"
	"fmt"
    "log"
    "gopkg.in/mgo.v2"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/services"
	"github.com/smart-evolution/smarthome/controllers"
	"github.com/smart-evolution/smarthome/controllers/api"
	"github.com/oskarszura/gowebserver"
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
    dbURI := os.Getenv("MONGOLAB_URI")
    addr, _ := getServerAddress()

    utils.VERSION = VERSION

    serverOptions := gowebserver.WebServerOptions{
        Port: addr,
        StaticFilesUrl: "/static/",
        StaticFilesDir: "public",
    }

    log.Println("Connecting to mgo with URI = " + dbURI)
    dbSession, err := mgo.Dial(dbURI)
    if err != nil {
        panic(err)
    }
    defer dbSession.Close()
    dbSession.SetMode(mgo.Monotonic, true)
    utils.SetSession(dbSession)

    services.InitInfluxService()
    go services.RunHomeService()

    server := gowebserver.New(serverOptions, controllers.NotFound)
    server.Router.AddRoute("/login/register", controllers.Register)
    server.Router.AddRoute("/login/logout", controllers.AuthenticateLogout)
    server.Router.AddRoute("/login", controllers.Authenticate)
    server.Router.AddRoute("/", controllers.CtrDashboard)
    server.Router.AddRoute("/agent/{agent}", controllers.CtrAgent)
    server.Router.AddRoute("/api/agents", api.CtrAgents)
    server.Router.AddRoute("/api/agents/{agent}", api.CtrAgents)
    server.Router.AddRoute("/api/alerts", api.CtrAlerts)
    server.Router.AddRoute("/api/sendalert", api.CtrSendAlert)

    server.Run()
}

