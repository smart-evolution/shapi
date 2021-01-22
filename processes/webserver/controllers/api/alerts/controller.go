package alerts

import (
	agentUsecases "github.com/smart-evolution/shapi/domain/usecases/agent"
	userUsecases "github.com/smart-evolution/shapi/domain/usecases/user"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/base"
)

// Controller -
type Controller struct {
	*base.Controller
	UserUsecases userUsecases.Usecase
}

// New -
func New(b *base.Controller, u userUsecases.Usecase, a agentUsecases.Usecase) *Controller {
	return &Controller{
		b,
		u,
	}
}
