package agents

import (
	agentUsecases "github.com/smart-evolution/shapi/domain/usecases/agent"
	userUsecases "github.com/smart-evolution/shapi/domain/usecases/user"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/base"
)

// Controller - agents controller
type Controller struct {
	*base.Controller
	AgentUsecases agentUsecases.Usecase
	UserUsecases  userUsecases.Usecase
}

// New - creates instance of post Controller
func New(b *base.Controller, au agentUsecases.Usecase, uu userUsecases.Usecase) *Controller {
	return &Controller{
		b,
		au,
		uu,
	}
}
