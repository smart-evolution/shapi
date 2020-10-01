package handlers

import (
	"github.com/coda-it/goutils/logger"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"net/http"
)

// Delete - handles agent delete action
func Delete(s state.IState, w http.ResponseWriter, href string, agentID string) {
	err := s.RemoveAgent(agentID)

	if err != nil {
		msg := "error deleting agent with ID = " + agentID
		logger.Log(msg)
		handlers.HandleError(w, href, msg, http.StatusInternalServerError)
	}
}
