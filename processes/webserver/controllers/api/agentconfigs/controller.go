package agentconfigs

import (
	agentUsecases "github.com/smart-evolution/shapi/domain/usecases/agent"
	userUsecases "github.com/smart-evolution/shapi/domain/usecases/user"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/base"
)

// Controller - agent configs controller
type Controller struct {
	*base.Controller
	UserUsecases  userUsecases.Usecase
	AgentUsecases agentUsecases.Usecase
}

// New - creates agent configs controller
func New(b *base.Controller, u userUsecases.Usecase, a agentUsecases.Usecase) *Controller {
	return &Controller{
		b,
		u,
		a,
	}
}
