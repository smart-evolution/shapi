package handlers

import (
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"github.com/smart-evolution/shapi/utils"
	"net/http"
)

// Delete - handles agent delete action
func Delete(s state.IState, w http.ResponseWriter, href string, agentID string) {
	err := s.RemoveAgent(agentID)

	if err != nil {
		msg := "error deleting agent with ID = " + agentID
		utils.Log(msg)
		handlers.HandleError(w, href, msg, http.StatusInternalServerError)
	}
}
