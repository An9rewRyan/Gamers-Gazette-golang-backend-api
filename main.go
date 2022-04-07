package main

import (
	"d/go/utils"
	"fmt"
	"net/http"
)

func main() {
	utils.Set_db(Db_conn_str)
	mux := http.NewServeMux()
	fmt.Println("Server is listening...")
	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	Set_urls(mux)
	server.ListenAndServe()
}
