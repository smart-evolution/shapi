package base

import (
	"encoding/json"
	"github.com/coda-it/goutils/logger"
	"github.com/coda-it/goutils/mailer"
	"github.com/coda-it/gowebserver/helpers"
	"net/http"
)

// Controller - base controller
type Controller struct {
	Mailer mailer.IMailer
}

// New - creates new instance of base Controller
func New(m mailer.IMailer) *Controller {
	return &Controller{
		m,
	}
}

// HandleErrorResponse - renders error output
func (c *Controller) HandleErrorResponse(w http.ResponseWriter, msg string) {
	logger.Log(msg)
	http.Error(w, msg, http.StatusInternalServerError)
}

// HandleJSONResponse - renders JSON output
func (c *Controller) HandleJSONResponse(w http.ResponseWriter, data interface{}, embedded interface{}, links map[string]map[string]string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))

	if err != nil {
		c.HandleErrorResponse(w, "error parsing JSON response")
	}
}
