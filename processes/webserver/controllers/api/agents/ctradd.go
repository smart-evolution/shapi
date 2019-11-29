package agents

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources"
	"github.com/smart-evolution/shapi/datasources/state"
	"github.com/smart-evolution/shapi/processes/webserver/handlers"
	utl "github.com/smart-evolution/shapi/utils"
	"io/ioutil"
	"net/http"
)

type body struct {
	Name string
	IP   string
	Type string
}

// CtrAdd - add new agent by IP
func CtrAdd(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
	handlers.CorsHeaders(w, r)

	st := s.GetDataSource(datasources.State)

	state, ok := st.(state.IState)

	if !ok {
		utl.Log("Store should implement IState")
		http.Error(w, "Store should implement IState", http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			utl.Log("error reading request body")
			http.Error(w, "error reading request body", http.StatusInternalServerError)
			return
		}

		var msg body
		err = json.Unmarshal(b, &msg)

		if err != nil {
			utl.Log("error parsing request body")
			http.Error(w, "error parsing request body", http.StatusInternalServerError)
			return
		}

		state.AddAgent("", msg.Name, msg.IP, msg.Type)

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
