package homebot

import (
    "log"
    "time"
    "sync"
    "github.com/influxdata/influxdb/client/v2"
    "github.com/smart-evolution/smarthome/models/agent"
    "github.com/smart-evolution/smarthome/datasources/dataflux"
    "github.com/smart-evolution/smarthome/datasources/state"
    "github.com/smart-evolution/smarthome/services/email"
)

// HomeBot - struct for homebot administrator
type HomeBot struct {
    store  dataflux.IDataFlux
    state  state.IState
    mailer email.IMailer
}

// New - creates new instances of HomeBot
func New(store dataflux.IDataFlux, mailer email.IMailer, st state.IState) *HomeBot {
    return &HomeBot {
        store: store,
        state: st,
        mailer: mailer,
    }
}

func persistData(store dataflux.IDataFlux) func(*agent.Agent, map[string]interface{}) {
    return func (agent *agent.Agent, data map[string]interface{}) {
        pt, _ := client.NewPoint(
            agent.ID(),
            map[string]string{ "home": agent.Name() },
            data,
            time.Now(),
        )

        err := store.AddData(pt)

        if err != nil {
            log.Println("homebot/persistData: ", err)
        }
    }
}

func (hb *HomeBot) runCommunicationLoop() {
    for range time.Tick(time.Second * 10) {
        if hb.store.IsConnected() == false {
            log.Println("homebot/runCommunicationLoop: cannot fetch packages, Influx is down")
            return
        }

        agents := hb.state.Agents()
        done := make(chan struct{})
        var wg sync.WaitGroup
        wg.Add(len(agents))

        for i := 0; i < len(agents); i++ {
            a := agents[i]
            log.Println("homebot/runCommunicationLoop: fetching from=", a.Name())

            if a.AgentType() == "type1" {
                go a.FetchPackage(hb.mailer.BulkEmail, persistData(hb.store), hb.state.IsAlerts(), &wg)
            }
        }

        go func() {
            wg.Wait()
            close(done)
        }()

        switch {
        case <-done:
            continue
        case <-time.After(3 * time.Second):
            continue
        }
    }
}

// RunService - setup and run everything
func (hb *HomeBot) RunService() {
    hb.runCommunicationLoop()
}


