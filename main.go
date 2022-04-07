package main

import (
	"d/go/utils"
	"fmt"
	"net/http"
)

func main() {
	utils.Set_db(Db_conn_str)
	Set_urls()
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8000", nil)
}
