package main

import (
	"aChat/src/data"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err = session(w, r)

		public_tmpl_files := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html"}

		private_tmpl_files := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html"}
		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		}
		templates.ExecuteTemplate(w, "layout", threads)
	}

}

func authenticate(w http.ResponseWriter, r *http.Request) {
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
