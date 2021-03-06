package routers

import (
	"d/go/views/basic_auth"
	"net/http"

	"github.com/gorilla/mux"
)

func Route_auth_basic(mux *mux.Router) {
	mux.HandleFunc("/checkname", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		basic_auth.Checkname(w, r)
	})
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		basic_auth.Signin(w, r)
	})
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		basic_auth.Signup(w, r)
	})
	mux.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		basic_auth.Refresh(w, r)
	})
	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		basic_auth.Logout(w, r)
	})
	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		basic_auth.Me(w, r)
	})
}
