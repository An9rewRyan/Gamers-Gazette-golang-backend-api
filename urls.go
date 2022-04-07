package main

import (
	"d/go/views"
	"net/http"
)

func Set_urls() {
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		views.About(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		views.Home(w, r)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
