package services

import (
    "log"
    "io/ioutil"
    "time"
    "strings"
    "github.com/influxdata/influxdb/client/v2"
    "github.com/smart-evolution/smarthome/models"
)

func persistData(agent models.Agent, data map[string]interface{}) {
    pt, _ := client.NewPoint(
        agent.ID,
        map[string]string{ "home": agent.Name },
        data,
        time.Now(),
    )

    InfluxBp.AddPoint(pt)
    err = InfluxClient.Write(InfluxBp)

    if err != nil {
        log.Println("services", err)
    }
}

func addAgent(id string, name string, url string, agentType string) {
    log.Println("services: adding home agent '" + name + "' with URL '" + url + "'")

    agent := models.Agent{
        ID: id,
        Name: name,
        URL: url,
        AgentType: agentType,
    }

    models.Agents = append(models.Agents, agent)
}

func setupAgents() {
    agentsCnf, err := ioutil.ReadFile("hardware/agents.config")

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

            addAgent(id, name, apiURL, agentType)
        }
    }
}

func runCommunicationLoop() {
    for range time.Tick(time.Second * 10) {
        if InfluxConnected == false {
            log.Println("services: cannot fetch packages, Influx is down")
            return
        }

        for i := 0; i < len(models.Agents); i++ {
            a := models.Agents[i]
            log.Println("services: fetching from=", a.Name)

            log.Println("---- agent: ", a. AgentType)
            if a.AgentType == "type1" {
                a.FetchPackage(BulkEmail, persistData)
            }
        }
    }
}

// RunHomeService - setup and run everything
func RunHomeService() {
    setupAgents()
    runCommunicationLoop()
}


