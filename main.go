package main

import (
	"os"
	"fmt"
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

func main() {
	addr, _ := getServerAddress()

serverOptions := gws.WebServerOptions{
	addr,
	"/static/",
	"public",
}

server := gws.New(serverOptions, controllers.NotFound)
server.Router.AddRoute("/", controllers.CtrDashboard)
server.Router.AddRoute("/api/home", api.CtrHome)

server.Run()
}

