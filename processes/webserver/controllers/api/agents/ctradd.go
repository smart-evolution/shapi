package agents

import (
	"encoding/json"
	"github.com/coda-it/gowebserver/helpers"
	"github.com/coda-it/gowebserver/router"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/store"
	"github.com/smart-evolution/shapi/datasources"
	"github.com/smart-evolution/shapi/datasources/state"
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

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

		json.NewEncoder(w).Encode(helpers.ServeHal(data, embedded, links))
	}
}
