package main

import (
	"os"
    "io/ioutil"
    "strings"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/datasources/persistence"
    "github.com/smart-evolution/smarthome/datasources/dataflux"
    "github.com/smart-evolution/smarthome/datasources/state"
    "github.com/smart-evolution/smarthome/processes/homebot"
    "github.com/smart-evolution/smarthome/services/email"
    "github.com/smart-evolution/smarthome/processes/webserver"
    "github.com/smart-evolution/smarthome/models/user"
    "github.com/smart-evolution/smarthome/models/agent"
    "gopkg.in/mgo.v2/bson"
)

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./version.go

func getAgents(hardwareFile string) []*agent.Agent {
    var agents []*agent.Agent
    agentsCnf, err := ioutil.ReadFile(hardwareFile)

    if err != nil {
        utils.Log(err)
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

            agents = append(agents, agent.New(id, name, apiURL, agentType))
        }
    }

    return agents
}

func getRecipients(p *persistence.Persistance) []string {
    var users []user.User
    var recipients []string

    c := p.GetCollection("users")
    err := c.Find(bson.M{}).All(&users)

    if err != nil {
        utils.Log("Alert recipients not found", err)
    }

    for _, u := range users {
        recipients = append(recipients, u.Username)
    }

    return recipients
}

func main() {
    utils.VERSION = VERSION

    agents := getAgents("hardware/agents.config")
    s := state.New(agents)

    p := persistence.New(
        os.Getenv("MONGOLAB_URI"),
        os.Getenv("DB_NAME"),
    )

    df := dataflux.New("http://localhost:8086")

    recipients := getRecipients(p)
    m := email.New(
        recipients,
        os.Getenv("EMAILNAME"),
        os.Getenv("EMAILPASS"),
        os.Getenv("SMTPPORT"),
        os.Getenv("SMTPAUTHURL"),
    )

    hb := homebot.New(df, p, m, s)
    go hb.RunService()

    ws := webserver.New(
        os.Getenv("PORT"),
        df,
        p,
        s,
    )
    ws.RunService()
}

