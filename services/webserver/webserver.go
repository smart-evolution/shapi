package webserver

import (
    "fmt"
    "github.com/coda-it/gowebserver"
    "github.com/smart-evolution/smarthome/services/webserver/controllers"
    "github.com/smart-evolution/smarthome/services/webserver/controllers/api"
    "github.com/smart-evolution/smarthome/interfaces"
)

// WebServer - adapter for gowebserver instance
type WebServer struct {
    server *gowebserver.WebServer
}

func getServerAddress(port string) (string, error) {
    if port == "" {
        return "", fmt.Errorf("Port not set")
    }
    return ":" + port, nil
}

// New - creates new WebServer instance
func New(port string, store interfaces.IDataFlux, persistence interfaces.IPersistance, s interfaces.IState) *WebServer {
    addr, _ := getServerAddress(port)
    serverOptions := gowebserver.WebServerOptions{
        Port: addr,
        StaticFilesUrl: "/static/",
        StaticFilesDir: "public",
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

    server.AddDataSource("dataflux", store)
    server.AddDataSource("persistence", persistence)
    server.AddDataSource("state", s)

    return &WebServer{
        server: server,
    }
}

// RunService - runs WebServer process
func (ws *WebServer) RunService() {
    ws.server.Run()
}
