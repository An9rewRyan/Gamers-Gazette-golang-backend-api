package main

import (
	"d/go/routers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	//uncomment on first launch and comment after sucess
	// database.Create_basic_tables()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1/").Subrouter()
	basic_auth := router.PathPrefix("/auth/").Subrouter()
	soc_auth := router.PathPrefix("/socialauth/").Subrouter()
	routers.Route_api(api)
	routers.Route_auth_basic(basic_auth)
	routers.Route_auth_social(soc_auth)

	fmt.Println("Server is listening....")

	handler := cors.Handler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Port: ", port)
	server := http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	server.ListenAndServe()
}
