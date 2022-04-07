package main

import (
	"d/go/views"
	"net/http"
)

func Set_urls(mux *http.ServeMux) {
	mux.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		views.About(w, r)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		} else {
			views.Home(w, r)
		}
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
