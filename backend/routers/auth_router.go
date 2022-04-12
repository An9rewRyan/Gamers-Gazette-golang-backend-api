package routers

import (
	"d/go/views/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func Route_auth(mux *mux.Router) {
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		auth.Signin(w, r)
	})
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		auth.Signup(w, r)
	})
	// mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
	// 	auth.Welcome(w, r)
	// })
	mux.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		auth.Refresh(w, r)
	})
	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.Logout(w, r)
	})
	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		auth.Me(w, r)
	})
	mux.HandleFunc("/oauth2_test", func(w http.ResponseWriter, r *http.Request) {
		auth.Oauth2(w, r)
	})
}
