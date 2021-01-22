package front

import "github.com/smart-evolution/shapi/processes/webserver/controllers/base"

// Controller - landing page controller
type Controller struct {
	*base.Controller
}

// New - creates new instance of landing page controller
func New(b *base.Controller) *Controller {
	return &Controller{
		b,
	}
}
