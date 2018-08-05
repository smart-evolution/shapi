package utils

import (
	"os"
	"log"
	"net/http"
	"path/filepath"
	"html/template"
    "github.com/oskarszura/smarthome/utils"
    "github.com/oskarszura/smarthome/services"
	. "github.com/oskarszura/smarthome/models"
	. "github.com/oskarszura/gowebserver/session"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, sm ISessionManager) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}

    menu := make([]string, len(services.Agents))
    for _, a := range services.Agents {
        menu = append(menu, a.Name)
    }

    params := make(map[string]interface{})
    params["menu"] = menu

    templateModel := Page{utils.VERSION, name, true, params}

	template := template.Must(
        template.ParseFiles(
            dir + "/views/" + name + ".html",
            dir + "/views/navigation.html",
            dir + "/views/view.html",
        ),
    )
	template.ExecuteTemplate(w, "base", templateModel)
}
