package front

import "github.com/smart-evolution/shapi/processes/webserver/controllers/base"

// Controller -
type Controller struct {
	*base.Controller
}

// New -
func New(b *base.Controller) *Controller {
	return &Controller{
		b,
	}
}
