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
    if(len(r.FormValue("emotion")) > 0){//only request/show ponies when an emotion was entered
        c := appengine.NewContext(r)
        client := urlfetch.Client(c)
        resp, _ := client.Get("http://ponyfac.es/api.json/tag:" + r.FormValue("emotion"))

        body, _ := ioutil.ReadAll(resp.Body)
        stable := new(Stable)
        json.Unmarshal(body, stable)

        for i := 0; i < len(stable.Faces); i++ {
            fmt.Fprintf(w, "<img src=\"%v\"/>", stable.Faces[i].Image)
        }
    }
}

const ponyForm = `
<html>
  <body>
    <form action="/handler" method="post">
      <div id="entry-form">
        <label for="emotion">Emotion:</label><input name="emotion" rows="3" cols="60">
        </input><input type="submit" value="Get Ponies">
        <div>
            <a href="http://www.cornify.com" onclick="cornify_add();return false;"><img src="http://www.cornify.com/assets/cornifycorn.gif" width="52" height="51" border="0" alt="Cornify" /></a><script type="text/javascript" src="http://www.cornify.com/js/cornify.js"></script>
        </div>
      </div>
    </form>
  </body>
</html>
`