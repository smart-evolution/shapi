package homebot

import (
    "log"
    "sync"
    "io/ioutil"
    "time"
    "strings"
    "errors"
    "github.com/influxdata/influxdb/client/v2"
    "github.com/smart-evolution/smarthome/models/agent"
    "github.com/smart-evolution/smarthome/state"
)

type HomeBot struct {
    hardwareFile string
    store state.IDataFlux
    mailer state.IMailer
    Agents []agent.Agent
}

func New(hardwareFile string, store state.IDataFlux, mailer state.IMailer) HomeBot {
    return HomeBot {
        hardwareFile: hardwareFile,
        store: store,
        mailer: mailer,
        Agents: []agent.Agent{},
    }
}

func (hb HomeBot) addAgent(id string, name string, url string, agentType string) {
    log.Println("services: adding home agent '" + name + "' with URL '" + url + "'")

    agent := agent.New(id, name, url, agentType)

    hb.Agents = append(hb.Agents, agent)
}

func (hb HomeBot) setupAgents() {
    agentsCnf, err := ioutil.ReadFile(hb.hardwareFile)

    if err != nil {
        log.Print("services", err)
    }

    agentsConf := strings.Split(string(agentsCnf), "\n")

    for _, c := range agentsConf {
        cnfRow := strings.Split(c, ":")

        if (len(cnfRow) == 4) {
            id := cnfRow[0]
            name := cnfRow[1]
            ip := cnfRow[2]
            agentType := cnfRow[3]
            apiURL := "http://" + ip + "/api"

            hb.addAgent(id, name, apiURL, agentType)
        }
    }
}

// FindAgentByID - find corresponding agent by ID
func (hb HomeBot) FindAgentByID(id string) (agent.Agent, error) {
    for _, a := range hb.Agents {
        if a.ID() == id {
            return a, nil
        }
    }

    return  agent.Agent{}, errors.New("No matching agent")
}

func persistData(store state.IDataFlux) func(agent.Agent, map[string]interface{}) {
    return func (agent agent.Agent, data map[string]interface{}) {
        pt, _ := client.NewPoint(
        agent.ID(),
        map[string]string{ "home": agent.Name() },
        data,
        time.Now(),
        )

        err := store.AddData(pt)

        if err != nil {
        log.Println("services", err)
        }
    }
}

func (hb HomeBot) runCommunicationLoop() {
    for range time.Tick(time.Second * 10) {
        if hb.store.IsConnected() == false {
            log.Println("services: cannot fetch packages, Influx is down")
            return
        }

        for i := 0; i < len(agent.Agents); i++ {
            a := agent.Agents[i]
            log.Println("services: fetching from=", a.Name)

            if a.AgentType() == "type1" {
                a.FetchPackage(hb.mailer.BulkEmail, persistData(hb.store), state.IsAlerts)
            }
        }
    }
}

// RunService - setup and run everything
func (hb HomeBot) RunService(wg sync.WaitGroup) {
    defer wg.Done()

    hb.setupAgents()
    hb.runCommunicationLoop()
}


