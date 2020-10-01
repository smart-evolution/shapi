package main

import (
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/goutils/mailer"
	"github.com/smart-evolution/shapi/datasources/dataflux"
	"github.com/smart-evolution/shapi/datasources/persistence"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/processes/cliserver"
	"github.com/smart-evolution/shapi/processes/homebot"
	"github.com/smart-evolution/shapi/processes/webserver"
	"github.com/smart-evolution/shapi/services/agentsniffer"
	"github.com/smart-evolution/shapi/utils"
	"gopkg.in/mgo.v2/bson"
	"os"
)

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./version.go

func getRecipients(p *persistence.Persistance) []string {
	var recipients []string

	users, err := p.FindAllUsers(bson.M{})

	if err != nil {
		logger.Log("alert recipients not found", err)
	}

	for _, u := range users {
		recipients = append(recipients, u.Username)
	}

	return recipients
}

func main() {
	SHPanelMongoURI := os.Getenv("SH_MONGO_URI")
	SHPanelMongoDB := os.Getenv("SH_MONGO_DB")
	SHHTTPPort := os.Getenv("SH_HTTP_PORT")

	logger.Log("Staring sh-api with the following ENV variables")
	logger.Log("SH_PANEL_MONGO_URI = " + SHPanelMongoURI)
	logger.Log("SH_PANEL_MONGO_DB = " + SHPanelMongoDB)
	logger.Log("SH_HTTP_PORT = " + SHHTTPPort)

	utils.VERSION = VERSION

	p := persistence.New(
		SHPanelMongoURI,
		SHPanelMongoDB,
	)

	s := state.New(p, []agent.IAgent{})

	df := dataflux.New(os.Getenv("SH_INFLUX_URI"))

	recipients := getRecipients(p)
	m := mailer.New(
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
		SHHTTPPort,
		df,
		p,
		s,
	)
	ws.RunService()
}
