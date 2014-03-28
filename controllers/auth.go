package controllers

import (
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/login", loginUser)
	http.HandleFunc("/logout", logoutUser)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, "/")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
}

func logoutUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusFound)
		return
	}

	url, err := user.LogoutURL(c, "/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
	return
}
