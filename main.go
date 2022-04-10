package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// utils.Set_db()
	// utils.Create_articles_table()
	// utils.Create_recently_loaded_articles_table()
	router := mux.NewRouter()
	fmt.Println("Server is listening....")
	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	Set_urls(router)
	// database.Create_test_articles()
	server.ListenAndServe()
}
