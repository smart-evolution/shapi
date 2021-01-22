package notfound

import (
	"github.com/smart-evolution/shapi/processes/webserver/controllers/base"
)

// Controller - 404 controller
type Controller struct {
	*base.Controller
}

// New - creates new instance of 404 controller
func New(b *base.Controller) *Controller {
	return &Controller{
		b,
	}
}
