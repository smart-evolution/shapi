package controllers

import (
    "net/http"
    "log"
    "github.com/smart-evolution/smarthome/processes/webserver/controllers/utils"
    "github.com/smart-evolution/smarthome/datasources/persistence"
    "github.com/smart-evolution/smarthome/models/user"
    "gopkg.in/mgo.v2/bson"
    "github.com/coda-it/gowebserver/router"
    "github.com/coda-it/gowebserver/session"
    "github.com/coda-it/gowebserver/store"
)

// Register - handle register page and register user process
func Register(w http.ResponseWriter, r *http.Request, opt router.UrlOptions, sm session.ISessionManager, s store.IStore) {
    switch r.Method {
    case "GET":
        utils.RenderTemplate(w, r, "register", sm, s)
    case "POST":
        var newUser *user.User

        dfc := s.GetDataSource("persistence")

        p, ok := dfc.(persistence.IPersistance);
        if !ok {
            log.Println("controllers: Invalid store ")
            return
        }

        c := p.GetCollection("users")

        newUser = &user.User{
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
