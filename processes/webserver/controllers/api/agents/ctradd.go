package agents

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	"github.com/smart-evolution/shapi/utils"
	"io/ioutil"
	"net/http"
)

type body struct {
	ID   string `json:"agentID"`
	Name string `json:"agentName"`
	IP   string `json:"agentIP"`
	Type string `json:"agentType"`
}

// CtrAdd - add new agent by IP
func CtrAdd(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	dst := s.GetDataSource(datasources.State)
	st, ok := dst.(state.IState)

	if !ok {
		utils.Log("Store should implement IState")
		http.Error(w, "Store should implement IState", http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			utils.Log("error reading request body")
			http.Error(w, "error reading request body", http.StatusInternalServerError)
			return
		}

		var msg body
		err = json.Unmarshal(b, &msg)

		if err != nil {
			utils.Log("error parsing request body")
			http.Error(w, "error parsing request body", http.StatusInternalServerError)
			return
		}

		st.AddAgent(msg.ID, msg.Name, msg.IP, msg.Type)

		data := struct {
			Devices string `json:"message"`
		}{
			"agent added",
		}

		embedded := map[string]string{}

		links := map[string]map[string]string{
			"self": map[string]string{
				"href": "/api/agents/add",
			},
		}

		handlers.HandleResponse(w, data, embedded, links, http.StatusOK)
	}
}
