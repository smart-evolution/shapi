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

// CorsHeaders - set required CORS headers
func (c *Controller) CorsHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

// HandleError - handle request which errored
func (c *Controller) HandleError(w http.ResponseWriter, href string, msg string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	data := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	embedded := map[string]string{}

	links := map[string]map[string]string{
		"self": map[string]string{
			"href": href,
		},
	}

	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))

	if err != nil {
		logger.Log("error parsing response")
	}
}
