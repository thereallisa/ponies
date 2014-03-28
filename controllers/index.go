package controllers

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/user"
	"appengine/urlfetch"
    "io/ioutil"
    "encoding/json"
)

func init() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/handler", handler)
}

type GlobalVariables struct {
	Title string
	User  *user.User
	PonyResults *Stable
}

// START FUNCTIONS FOR INDEX PAGE

var indexTemplate = template.Must(template.New("index.html").ParseFiles(
	"views/index.html",
	"views/layout/bootstrap_base_head.html",
	"views/layout/navbar.html",
	"views/layout/bootstrap_base_foot.html",
))

type IndexPageVariables struct {
	Global GlobalVariables
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	page_variables := IndexPageVariables{}

	page_variables.Global.Title = "Simply Ninja"
	page_variables.Global.User = user.Current(c)

	if err := indexTemplate.Execute(w, page_variables); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// END FUNCTIONS FOR INDEX PAGE


type Stable struct{
    Faces []Pony
}
type Pony struct{
    Image string
}

func handler(w http.ResponseWriter, r *http.Request) {

    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, _ := client.Get("http://ponyfac.es/api.json/tag:" + r.FormValue("emotion"))

    body, _ := ioutil.ReadAll(resp.Body)
    stable := new(Stable)
    json.Unmarshal(body, stable)

	page_variables := IndexPageVariables{}

	page_variables.Global.PonyResults = stable

	if err := indexTemplate.Execute(w, page_variables); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
