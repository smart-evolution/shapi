package main

import (
	"os"
	"fmt"
    "github.com/oskarszura/smarthome/utils"
    "github.com/oskarszura/smarthome/services"
	"github.com/oskarszura/smarthome/controllers"
	"github.com/oskarszura/smarthome/controllers/api"
	gws "github.com/oskarszura/gowebserver"
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

    serverOptions := gws.WebServerOptions{
        addr,
        "/static/",
        "public",
    }

    services.InitInfluxService()
    go services.RunHomeService()

    server := gws.New(serverOptions, controllers.NotFound)
    server.Router.AddRoute("/", controllers.CtrDashboard)
    server.Router.AddRoute("/api/home", api.CtrHome)
    server.Router.AddRoute("/api/alerts", api.CtrAlerts)
    server.Router.AddRoute("/api/sendalert", api.CtrSendAlert)

    server.Run()
}

