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

	go cliserver.RunService("3333")

	go agentsniffer.SniffAgents(s)

	ws := webserver.New(
		os.Getenv("PORT"),
		df,
		p,
		s,
	)
	ws.RunService()
}
