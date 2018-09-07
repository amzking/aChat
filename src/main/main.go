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
	mux.Handler("/authenticate", authenticate)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
