package handlers

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"github.com/smart-evolution/shapi/utils"
	"net/http"
)

// HandleError - handle request which errored
func HandleError(w http.ResponseWriter, href string, msg string, status int) {
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
		utils.Log("error parsing response")
	}
}
