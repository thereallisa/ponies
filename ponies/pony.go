package ponies

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io/ioutil"

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
	fmt.Fprint(w, ponyForm)
}

func handleEmo(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, _ := client.Get("http://ponyfac.es/api.json/tag:" + r.FormValue("emotion"))

	body, _ := ioutil.ReadAll(resp.Body)
	stable := new(Stable)
	json.Unmarshal(body, stable)

	fmt.Fprint(w, ponyForm)
	fmt.Fprintf(w, "<html>")

	for i := 0; i < len(stable.Faces); i++ {
		fmt.Fprintf(w, "<img src=\"%v\"/>", stable.Faces[i].Image)
	}
	fmt.Fprintf(w, "</html>")
}

const ponyForm = `
<html>
<head>
<link rel="icon" href="../favicon.ico"/>
<title>Emo Ponies</title>
</head>
  <body>
    <form action="/handleEmo" method="post">
      <div id="entry-form">
        <label for="emotion">Emotion:</label><input name="emotion" rows="3" cols="60"/><input type="submit" value="Get Ponies">
        <div>
            <a href="http://www.cornify.com" onclick="cornify_add();return false;"><img src="http://www.cornify.com/assets/cornifycorn.gif" width="52" height="51" border="0" alt="Cornify" /></a><script type="text/javascript" src="http://www.cornify.com/js/cornify.js"></script>
        </div>
      </div>
    </form>
  </body>
</html>
`
