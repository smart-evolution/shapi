package handlers

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"net/http"
)

func HandleError(w http.ResponseWriter, href string, msg string, status int) {
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
	json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
}
