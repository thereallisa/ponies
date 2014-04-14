package ponies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"appengine"
	"appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/handleEmo", handleEmo)
}

type Stable struct {
	Faces []Pony
}
type Pony struct {
	Image string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("ponies/layout.html"))
	t.Execute(w, map[string]string{"content": ""})
}

func handleEmo(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, _ := client.Get("http://ponyfac.es/api.json/tag:" + r.FormValue("emotion"))

	body, _ := ioutil.ReadAll(resp.Body)
	stable := new(Stable)
	json.Unmarshal(body, stable)

	var content bytes.Buffer
	for _, pony := range stable.Faces {
		content.WriteString(fmt.Sprintf("<img src=\"%v\"/>", pony.Image))
	}
	t := template.Must(template.ParseFiles("ponies/layout.html"))
	t.Execute(w, map[string]string{"content": content.String()})
}
