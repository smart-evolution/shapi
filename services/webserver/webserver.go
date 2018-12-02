package webserver

import (
    "fmt"
    "sync"
    "github.com/coda-it/gowebserver"
    "github.com/smart-evolution/smarthome/services/webserver/controllers"
    "github.com/smart-evolution/smarthome/services/webserver/controllers/api"
    "github.com/smart-evolution/smarthome/state"
)

type WebServer struct {
    server *gowebserver.WebServer
    store state.IDataFlux
}

func (ws WebServer) Store() state.IDataFlux {
    return ws.store
}

func getServerAddress(port string) (string, error) {
    if port == "" {
        return "", fmt.Errorf("$PORT not set")
    }
    return ":" + port, nil
}

func New(port string, store state.IDataFlux) WebServer {
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

    server.AddDataSource("datainflux", store)

    return WebServer{
        server: server,
        store: store,
    }
}

func (ws WebServer) RunService(wg sync.WaitGroup) {
    defer wg.Done()
    ws.server.Run()
}
