package resetdb

import (
	platformUsecases "github.com/smart-evolution/shapi/domain/usecases/platform"
	"github.com/smart-evolution/shapi/processes/webserver/controllers/base"
)

// Controller -
type Controller struct {
	*base.Controller
	PlatformUsecases platformUsecases.Usecase
}

// New -
func New(b *base.Controller, pu platformUsecases.Usecase) *Controller {
	return &Controller{
		b,
		pu,
	}
}
