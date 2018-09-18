package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("config.static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// 路由
	mux.HandleFunc("/", index)

	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)

	mux.HandleFunc("/authenticate", authenticate)

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
