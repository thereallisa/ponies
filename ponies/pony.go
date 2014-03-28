package ponies

import (
    "fmt"
    "net/http"

    "appengine"
    "appengine/urlfetch"
    "io/ioutil"
    "encoding/json"
)

func init() {
    http.HandleFunc("/", handler)
}
type Stable struct{
    Faces []Pony
}
type Pony struct{
    Image string
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, ponyForm)

    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, _ := client.Get("http://ponyfac.es/api.json/tag:" + r.FormValue("emotion"))

    body, _ := ioutil.ReadAll(resp.Body)
    stable := new(Stable)
    json.Unmarshal(body, stable)

    for i := 0; i < len(stable.Faces); i++ {
        fmt.Fprintf(w, "<html><img src=\"%v\"/></html>", stable.Faces[i].Image)
    }
}

const ponyForm = `
<html>
  <body>
    <form action="/handler" method="post">
      <div><input name="emotion" rows="3" cols="60"></input></div>
      <div><input type="submit" value="Request Pony Emotion"></div>
    </form>
  </body>
</html>`