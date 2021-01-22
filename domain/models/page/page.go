package page

// Page - entity representing page
type Page struct {
	Version  string
	Title    string
	IsLogged bool
	Params   map[string]interface{}
}
