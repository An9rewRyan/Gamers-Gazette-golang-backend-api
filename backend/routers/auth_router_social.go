package routers

import (
	"d/go/views/social_auth"
	"net/http"

	"github.com/gorilla/mux"
)

func Route_auth_social(mux *mux.Router) {
	mux.HandleFunc("/vk", func(w http.ResponseWriter, r *http.Request) {
		social_auth.Vk(w, r)
	})
}
