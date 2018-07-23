package utils

import (
	"os"
	"log"
	"net/http"
	"path/filepath"
	"html/template"
	. "github.com/oskarszura/smarthome/models"
	. "github.com/oskarszura/gowebserver/session"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, sm ISessionManager) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}

	template := template.Must(template.ParseFiles(dir + "/views/" + name + ".html",
		dir + "/views/navigation.html", dir + "/views/view.html"))
	templateModel := Page{VERSION, name, true}
	template.ExecuteTemplate(w, "base", templateModel)
}
