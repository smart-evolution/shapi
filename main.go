package main

import (
	"github.com/smart-evolution/smarthome/datasources/dataflux"
	"github.com/smart-evolution/smarthome/datasources/persistence"
	"github.com/smart-evolution/smarthome/datasources/state"
	"github.com/smart-evolution/smarthome/models/agent"
	"github.com/smart-evolution/smarthome/models/user"
	"github.com/smart-evolution/smarthome/processes/cliserver"
	"github.com/smart-evolution/smarthome/processes/homebot"
	"github.com/smart-evolution/smarthome/processes/webserver"
	"github.com/smart-evolution/smarthome/services/agentsniffer"
	"github.com/smart-evolution/smarthome/services/email"
	"github.com/smart-evolution/smarthome/utils"
	"gopkg.in/mgo.v2/bson"
	"os"
)

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./version.go

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

	s := state.New([]agent.IAgent{})

	p := persistence.New(
		os.Getenv("SH_MONGO_URI"),
		os.Getenv("SH_MONGO_DB"),
	)

	df := dataflux.New(os.Getenv("SH_INFLUX_URI"))

	recipients := getRecipients(p)
	m := email.New(
		recipients,
		os.Getenv("SH_MAILER_EMAIL_NAME"),
		os.Getenv("SH_MAILER_EMAIL_PASS"),
		os.Getenv("SH_MAILER_SMTP_PORT"),
		os.Getenv("SH_MAILER_SMTP_AUTHURL"),
	)

	hb := homebot.New(df, p, m, s)
	go hb.RunService()

	go cliserver.RunService(os.Getenv("SH_CLI_TCP_PORT"))

	go agentsniffer.SniffAgents(s)

	ws := webserver.New(
		os.Getenv("SH_HTTP_PORT"),
		df,
		p,
		s,
	)
	ws.RunService()
}
