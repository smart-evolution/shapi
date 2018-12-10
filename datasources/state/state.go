package state

import (
    "log"
    "errors"
    "io/ioutil"
    "strings"
    "github.com/smart-evolution/smarthome/models/agent"
)

// State
type State struct {
    isAlerts    bool
    sendAlert   bool
    agents      []agent.Agent
}

func New() State {
    return State{
        isAlerts: false,
        sendAlert: false,
        agents: []agent.Agent{},
    }
}

func (s State) Agents() []agent.Agent {
    return s.agents
}

func (s State) SetIsAlerts(i bool) {
    s.isAlerts = i
}

func (s State) IsAlerts() bool {
    return s.isAlerts
}

func (s State) SetSendAlert(i bool) {
    s.sendAlert = i
}

func (s State) SendAlert() bool {
    return s.sendAlert
}

func (s State) addAgent(id string, name string, url string, agentType string) {
    log.Println("services: adding home agent '" + name + "' with URL '" + url + "'")

    agent := agent.New(id, name, url, agentType)

    s.agents = append(s.agents, agent)
}

func (s State) SetupAgents(hardwareFile string) {
    agentsCnf, err := ioutil.ReadFile(hardwareFile)

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

            s.addAgent(id, name, apiURL, agentType)
        }
    }
}

// FindAgentByID - find corresponding agent by ID
func (s State) FindAgentByID(id string) (agent.Agent, error) {
    for _, a := range s.agents {
        if a.ID() == id {
            return a, nil
        }
    }

    return  agent.Agent{}, errors.New("No matching agent")
}
