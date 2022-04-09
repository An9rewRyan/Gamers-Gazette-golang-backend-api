package main

import (
	"d/go/views"
	"d/go/views/api"
	"net/http"

	"github.com/gorilla/mux"
)

func Set_urls(mux *mux.Router) {
	mux.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		views.About(w, r)
	})
	mux.HandleFunc("/api/articles/", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticles(w, r)
	})
	mux.HandleFunc("/api/articles/{id}/", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticleGet(w, r)
	})
	mux.HandleFunc("/api/articles/{id}/delete/", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticleDelete(w, r)
	})
	mux.HandleFunc("/api/articles/create", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticleCreate(w, r)
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
