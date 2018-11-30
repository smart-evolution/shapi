package homebot

import (
    "log"
    "io/ioutil"
    "time"
    "strings"
    "github.com/influxdata/influxdb/client/v2"
    "github.com/smart-evolution/smarthome/models/agent"
    "github.com/smart-evolution/smarthome/utils"
)

type HomeBot struct {
    HardwareFile    string
    Agents          []agent.Agent
}

func New(hardwareFile string) HomeBot {
    return HomeBot {
        HardwareFile: hardwareFile,
        Agents: []agent.Agent{},
    }
}

func (hb *HomeBot) addAgent(id string, name string, url string, agentType string) {
    log.Println("services: adding home agent '" + name + "' with URL '" + url + "'")

    agent := agent.Agent{
        ID: id,
        Name: name,
        URL: url,
        AgentType: agentType,
    }

    hb.Agents = append(hb.Agents, agent)
}

func (hb *HomeBot) setupAgents() {
    agentsCnf, err := ioutil.ReadFile(hb.HardwareFile)

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

func persistData(agent agent.Agent, data map[string]interface{}) {
    pt, _ := client.NewPoint(
        agent.ID,
        map[string]string{ "home": agent.Name },
        data,
        time.Now(),
    )

    utils.DataFlux.BatchPoints.AddPoint(pt)
    err := utils.DataFlux.Client.Write(utils.DataFlux.BatchPoints)

    if err != nil {
        log.Println("services", err)
    }
}

func (hb *HomeBot) runCommunicationLoop() {
    for range time.Tick(time.Second * 10) {
        if utils.DataFlux.IsConnected == false {
            log.Println("services: cannot fetch packages, Influx is down")
            return
        }

        for i := 0; i < len(agent.Agents); i++ {
            a := agent.Agents[i]
            log.Println("services: fetching from=", a.Name)

            if a.AgentType == "type1" {
                a.FetchPackage(utils.Mailer.BulkEmail, persistData)
            }
        }
    }
}

// RunHomeService - setup and run everything
func (hb *HomeBot) RunHomeService() {
    hb.setupAgents()
    hb.runCommunicationLoop()
}


