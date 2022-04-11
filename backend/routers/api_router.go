package routers

import (
	"d/go/views/api"
	"net/http"

	"github.com/gorilla/mux"
)

func Route_api(mux *mux.Router) {
	mux.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticles(w, r)
	})
	mux.HandleFunc("/articles/{id}/", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticleGet(w, r)
	})
	mux.HandleFunc("/articles/{id}/delete/", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticleDelete(w, r)
	})
	mux.HandleFunc("/articles/create", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticleCreate(w, r)
	})
	mux.HandleFunc("/articles/{id}/update", func(w http.ResponseWriter, r *http.Request) {
		api.ApiArticleUpdate(w, r)
	})
}
