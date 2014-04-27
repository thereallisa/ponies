package ponies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
type Emotions struct {
    Tags []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	emotions := getEmotions(w, r)
	t := template.Must(template.ParseFiles("ponies/layout.html"))
    t.Execute(w, map[string]string{"emotionContent": emotions})
}

func getEmotions(w http.ResponseWriter, r *http.Request) (string) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	emoResp, _ := client.Get("http://ponyfac.es/api.json/tags")

	body, _ := ioutil.ReadAll(emoResp.Body)
	emotions := new(Emotions)
	json.Unmarshal(body, emotions)

	var emotionContent bytes.Buffer
	for _, emotion := range emotions.Tags {
		emotion = strings.Replace(emotion,"%","",-1);
		selected := ""
		if(r.FormValue("emotion") == emotion) {
			selected = "selected"
		}
		emotionContent.WriteString(fmt.Sprintf("<option value=\"%v\" charset=\"UTF-8\" %v>%v</option>", emotion, selected, emotion))
	}	

	return emotionContent.String()
}

func handleEmo(w http.ResponseWriter, r *http.Request) {
	emotions := getEmotions(w, r)
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
	t.Execute(w, map[string]string{"content": content.String(), "emotionContent": emotions})	
}
