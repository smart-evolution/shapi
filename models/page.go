package models

type Page struct {
    Version     string
    Title       string
    IsLogged    bool
    Params      map[string]interface{}
}
