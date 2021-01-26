package main

import (
	"github.com/coda-it/goappframe"
	"github.com/coda-it/goappframe/module"
	"github.com/coda-it/goappframe/route"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/goutils/mailer"
	"github.com/smart-evolution/shapi/constants"
	"github.com/smart-evolution/shapi/data/dataflux"
	"github.com/smart-evolution/shapi/data/persistence"
	agentRepository "github.com/smart-evolution/shapi/data/repositories/agent"
	agentConfigsRepository "github.com/smart-evolution/shapi/data/repositories/agentconfigs"
	platformRepository "github.com/smart-evolution/shapi/data/repositories/platform"
	stateRepository "github.com/smart-evolution/shapi/data/repositories/state"
	userRepository "github.com/smart-evolution/shapi/data/repositories/user"
	"github.com/smart-evolution/shapi/domain/models/agent"
	agentUsecases "github.com/smart-evolution/shapi/domain/usecases/agent"
	platformUsecases "github.com/smart-evolution/shapi/domain/usecases/platform"
	userUsecases "github.com/smart-evolution/shapi/domain/usecases/user"
	"github.com/smart-evolution/shapi/processes/cliserver"
	"github.com/smart-evolution/shapi/processes/homebot"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/agentconfigs"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/agents"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/alerts"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/front"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/login"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/notfound"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/resetdb"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/sendalert"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/api/sniffagents"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/base"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/register"
	"github.com/smart-evolution/shapi/utils"
	"os"
)

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./version.go

func getRecipients(uu *userUsecases.Usecase) []string {
	var recipients []string

	users, err := uu.FindAllUsers()

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
	df := dataflux.New(os.Getenv("SH_INFLUX_URI"))

	stateRepositoryEntity := stateRepository.New(p, []agent.IAgent{})
	userRepositoryEntity := userRepository.New(p)
	agentRepositoryEntity := agentRepository.New(df)
	agentConfigsRepositoryEntity := agentConfigsRepository.New(p)
	platformRepositoryEntity := platformRepository.New(p)

	agentUsecasesEntity := agentUsecases.New(stateRepositoryEntity, agentRepositoryEntity, agentConfigsRepositoryEntity)
	userUsecasesEntity := userUsecases.New(stateRepositoryEntity, userRepositoryEntity)
	platformUsecasesEntity := platformUsecases.New(platformRepositoryEntity)

	baseController := base.New(mailer.New(
		[]string{},
		os.Getenv("WEBAPP_MAILER_EMAIL_NAME"),
		os.Getenv("WEBAPP_MAILER_EMAIL_PASS"),
		os.Getenv("WEBAPP_MAILER_SMTP_PORT"),
		os.Getenv("WEBAPP_MAILER_SMTP_AUTHURL"),
	))

	agentsCtl := agents.New(baseController, *agentUsecasesEntity, *userUsecasesEntity)
	apiAgentsModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/api/agents/add",
				Method:    "POST",
				Handler:   agentsCtl.CtrAdd,
				Protected: false,
			},
			{
				Path:      "/api/agents/{agent}",
				Method:    "OPTIONS",
				Handler:   agentsCtl.CtrAgentsOptions,
				Protected: false,
			},
			{
				Path:      "/api/agents/{agent}",
				Method:    "GET",
				Handler:   agentsCtl.CtrAgentsGet,
				Protected: false,
			},
			{
				Path:      "/api/agents/{agent}",
				Method:    "DELETE",
				Handler:   agentsCtl.CtrAgentsDelete,
				Protected: false,
			},
			{
				Path:      "/api/agents/{agent}",
				Method:    "POST",
				Handler:   agentsCtl.CtrAgentsPost,
				Protected: false,
			},
		},
	}

	agentConfigsCtl := agentconfigs.New(baseController, *userUsecasesEntity, *agentUsecasesEntity)
	apiAgentConfigsModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/api/agent-configs/{agent}",
				Method:    "OPTIONS",
				Handler:   agentConfigsCtl.CtrAgentConfigOptions,
				Protected: false,
			},
			{
				Path:      "/api/agent-configs/{agent}",
				Method:    "GET",
				Handler:   agentConfigsCtl.CtrAgentConfigGet,
				Protected: false,
			},
			{
				Path:      "/api/agent-configs/{agent}",
				Method:    "POST",
				Handler:   agentConfigsCtl.CtrAgentConfigPost,
				Protected: false,
			},
		},
	}

	userRegisterCtl := register.New(baseController, *userUsecasesEntity)
	userRegisterModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/login/register",
				Method:    "POST",
				Handler:   userRegisterCtl.CtrRegisterPost,
				Protected: false,
			},
		},
	}

	userLoginCtl := login.New(baseController, *userUsecasesEntity)
	userLoginModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/api/login",
				Method:    "OPTIONS",
				Handler:   userLoginCtl.CtrLoginOptions,
				Protected: false,
			},
			{
				Path:      "/api/login",
				Method:    "POST",
				Handler:   userLoginCtl.CtrLoginPost,
				Protected: false,
			},
		},
	}

	resetDbCtl := resetdb.New(baseController, *platformUsecasesEntity)
	platformModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/api/reset",
				Method:    "POST",
				Handler:   resetDbCtl.CtrResetDbPost,
				Protected: false,
			},
		},
	}

	frontCtl := front.New(baseController)
	frontModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/api",
				Method:    "ALL",
				Handler:   frontCtl.CtrFrontAll,
				Protected: false,
			},
		},
	}

	alertsCtl := alerts.New(baseController, *userUsecasesEntity, *agentUsecasesEntity)
	alertsModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/api/alerts",
				Method:    "GET",
				Handler:   alertsCtl.CtrAlertsGet,
				Protected: false,
			},
			{
				Path:      "/api/alerts",
				Method:    "OPTIONS",
				Handler:   alertsCtl.CtrAlertsOptions,
				Protected: false,
			},
			{
				Path:      "/api/alerts",
				Method:    "POST",
				Handler:   alertsCtl.CtrAlertsPost,
				Protected: false,
			},
		},
	}

	sendalertCtl := sendalert.New(baseController, *userUsecasesEntity)
	sendalertModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/api/sendalert",
				Method:    "ALL",
				Handler:   sendalertCtl.CtrSendAlertAll,
				Protected: false,
			},
		},
	}

	sniffagentsCtl := sniffagents.New(baseController, *agentUsecasesEntity)
	sniffagentsModule := module.Module{
		Enabled: true,
		Routes: []route.Route{
			{
				Path:      "/api/sniffagents",
				Method:    "ALL",
				Handler:   sniffagentsCtl.CtrSniffAgentsAll,
				Protected: false,
			},
		},
	}

	notFoundCtl := notfound.New(baseController)
	webApp := goappframe.New(goappframe.Internals{
		Port: SHHTTPPort,
		Modules: []module.Module{
			apiAgentsModule,
			apiAgentConfigsModule,
			userRegisterModule,
			userLoginModule,
			platformModule,
			frontModule,
			alertsModule,
			sendalertModule,
			sniffagentsModule,
		},
		Persistence: p,
		DataKey:     constants.PersistenceDataKey,
		Mailer: mailer.New(
			[]string{},
			os.Getenv("WEBAPP_MAILER_EMAIL_NAME"),
			os.Getenv("WEBAPP_MAILER_EMAIL_PASS"),
			os.Getenv("WEBAPP_MAILER_SMTP_PORT"),
			os.Getenv("WEBAPP_MAILER_SMTP_AUTHURL"),
		),
		NotFound: notFoundCtl.CtrNotFound,
	})

	recipients := getRecipients(userUsecasesEntity)
	m := mailer.New(
		recipients,
		os.Getenv("SH_MAILER_EMAIL_NAME"),
		os.Getenv("SH_MAILER_EMAIL_PASS"),
		os.Getenv("SH_MAILER_SMTP_PORT"),
		os.Getenv("SH_MAILER_SMTP_AUTHURL"),
	)

	hb := homebot.New(df, *agentConfigsRepositoryEntity, m, *stateRepositoryEntity)
	go hb.RunService()

	go cliserver.RunService(os.Getenv("SH_CLI_TCP_PORT"))

	go agentUsecasesEntity.SniffAgents()

	webApp.Run()
}
