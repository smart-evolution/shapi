package utils

import (
	"os"
	"net/http"
	"path/filepath"
	"html/template"
    utl "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/datasources/state"
	"github.com/smart-evolution/smarthome/models/agent"
    "github.com/smart-evolution/smarthome/models/page"
	"github.com/coda-it/gowebserver/session"
    "github.com/coda-it/gowebserver/store"
)


// RenderTemplate - helper for page rendering
func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, sm session.ISessionManager, s store.IStore) {
    sessionID, _ := GetSessionID(r)
    isLogged := sm.IsExist(sessionID)
    isPrivate := IsRequestFromIntranet(r)

    if !isLogged {
        ClearSession(w)

        if r.URL.Path != "/login" && r.URL.Path != "/login/register" {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
        }
    }

    if !isPrivate && r.URL.Path == "/login/register" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }

    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

    if err != nil {
        utl.Log(err)
    }

    st := s.GetDataSource("state")

    state, ok := st.(state.IState);
    if !ok {
        utl.Log("Invalid store")
        return
    }

    menu := make([]*agent.Agent, 0)
    for _, a := range state.Agents() {
        menu = append(menu, a)
    }

    params := make(map[string]interface{})
    params["menu"] = menu

    templateModel := page.Page{
        Version: utils.VERSION,
        Title: name,
        IsLogged: isLogged,
        Params: params,
    }

	template := template.Must(
        template.ParseFiles(
            dir + "/views/" + name + ".html",
            dir + "/views/navigation.html",
            dir + "/views/view.html",
        ),
    )
	template.ExecuteTemplate(w, "base", templateModel)
}
