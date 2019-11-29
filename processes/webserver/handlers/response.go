package handlers

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"net/http"
)

// HandleResponse - handle response with proper headers
func HandleResponse(w http.ResponseWriter, data interface{}, embedded interface{}, links map[string]map[string]string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
}
