package main

import (
	"d/go/views"
	"net/http"

	"github.com/gorilla/mux"
)

func Set_urls(mux *mux.Router) {
	mux.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		views.About(w, r)
	})
	mux.HandleFunc("/api/articles/", func(w http.ResponseWriter, r *http.Request) {
		views.ApiArticles(w, r)
	})
	mux.HandleFunc("/api/articles/{id}/", func(w http.ResponseWriter, r *http.Request) {
		views.ApiArticleGet(w, r)
	})
	mux.HandleFunc("/api/articles/{id}/delete/", func(w http.ResponseWriter, r *http.Request) {
		views.ApiArticleDelete(w, r)
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
