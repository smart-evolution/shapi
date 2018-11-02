package utils

import (
	"os"
	"log"
	"net/http"
	"path/filepath"
	"html/template"
    "github.com/oskarszura/smarthome/utils"
    "github.com/oskarszura/smarthome/services"
	"github.com/oskarszura/smarthome/models"
	"github.com/oskarszura/gowebserver/session"
)

// RenderTemplate - helper for page rendering
func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, sm session.ISessionManager) {
    sessionID := utils.GetSessionID(r)
    isLogged := sm.IsExist(sessionID)

    if !isLogged && r.URL.Path != "/login" && r.URL.Path != "/login/register" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}

    menu := make([]services.Agent, 0)
    for _, a := range services.Agents {
        menu = append(menu, a)
    }

    params := make(map[string]interface{})
    params["menu"] = menu

    templateModel := models.Page{utils.VERSION, name, true, params}

	template := template.Must(
        template.ParseFiles(
            dir + "/views/" + name + ".html",
            dir + "/views/navigation.html",
            dir + "/views/view.html",
        ),
    )
	template.ExecuteTemplate(w, "base", templateModel)
}
