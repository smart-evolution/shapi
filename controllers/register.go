package controllers

import (
    "net/http"
    "log"
    "github.com/oskarszura/smarthome/utils"
    ctrutl "github.com/oskarszura/smarthome/controllers/utils"
    "github.com/oskarszura/smarthome/models"
    "gopkg.in/mgo.v2/bson"
    "github.com/oskarszura/gowebserver/router"
    "github.com/oskarszura/gowebserver/session"
)

func Register(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    switch r.Method {
    case "GET":
        ctrutl.RenderTemplate(w, r, "register", sm)
    case "POST":
        var newUser *models.User

        ds := utils.GetDataSource()
        c := ds.C("users")

        newUser = &models.User{
            Id: bson.NewObjectId(),
            Username: r.PostFormValue("username"),
            Password: r.PostFormValue("password"),
        }

        err := c.Insert(newUser)
        if err != nil {
            log.Fatalln(err)
        }

        log.Println("Registered user", newUser)

        http.Redirect(w, r, "/", http.StatusSeeOther)
    default:
    }
}
