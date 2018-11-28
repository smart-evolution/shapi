package controllers

import (
    "net/http"
    "log"
    "github.com/smart-evolution/smarthome/utils"
    ctrutl "github.com/smart-evolution/smarthome/controllers/utils"
    "github.com/smart-evolution/smarthome/models"
    "gopkg.in/mgo.v2/bson"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
)

// Register - handle register page and register user process
func Register(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager) {
    switch r.Method {
    case "GET":
        ctrutl.RenderTemplate(w, r, "register", sm)
    case "POST":
        var newUser *models.User

        ds := utils.GetDataSource()
        c := ds.C("users")

        newUser = &models.User{
            ID: bson.NewObjectId(),
            Username: r.PostFormValue("username"),
            Password: utils.HashString(r.PostFormValue("password")),
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
