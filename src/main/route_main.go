package main

import (
	"aChat/src/data"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err == nil {
		return
	}

	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, _ := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
