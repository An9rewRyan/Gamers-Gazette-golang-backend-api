package main

import (
	// "d/go/utils/database"
	"fmt"
	"net/http"

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
	fmt.Println("Server is listening....")
	handler := cors.Handler(router)
	Set_urls(router)
	server := http.Server{
		Addr:    ":8000",
		Handler: handler,
	}

	// database.Create_test_articles()
	server.ListenAndServe()
}
