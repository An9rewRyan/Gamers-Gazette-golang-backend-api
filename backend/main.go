package main

import (
	"d/go/routers"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	//uncomment on first launch and comment after sucess
	// database.Create_basic_tables()
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println("Path to exec: ", exPath)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1/").Subrouter()
	auth := router.PathPrefix("/auth/").Subrouter()
	routers.Route_api(api)
	routers.Route_auth(auth)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../web")))
	router.PathPrefix("/articles/").Handler(http.FileServer(http.Dir("../web")))
	fmt.Println("Server is listening....")

	handler := cors.Handler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Port: ", port)
	server := http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	// database.Create_test_articles()
	server.ListenAndServe()
}
