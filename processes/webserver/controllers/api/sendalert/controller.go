package sendalert

import (
	userUsecases "github.com/smart-evolution/shapi/domain/usecases/user"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/base"
)

// Controller - send alert controller
type Controller struct {
	*base.Controller
	UserUsecases userUsecases.Usecase
}

// New - creates new instance of send alert controller
func New(b *base.Controller, uu userUsecases.Usecase) *Controller {
	return &Controller{
		b,
		uu,
	}
}
