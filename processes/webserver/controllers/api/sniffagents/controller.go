package sniffagents

import (
	agentUsecases "github.com/smart-evolution/shapi/domain/usecases/agent"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/base"
)

// Controller -
type Controller struct {
	*base.Controller
	UserUsecases agentUsecases.Usecase
}

// New -
func New(b *base.Controller, au agentUsecases.Usecase) *Controller {
	return &Controller{
		b,
		au,
	}
}
