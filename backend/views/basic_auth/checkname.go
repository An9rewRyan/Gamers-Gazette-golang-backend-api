package basic_auth

import (
	"context"
	"d/go/structs"
	"d/go/utils/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jackc/pgx/v4"
)

func Checkname(w http.ResponseWriter, r *http.Request) {
	storedCreds := &structs.Credentials{}
	db, err := database.Connect_db()
	if err != nil {
		fmt.Println("err")
		return
	}
	defer db.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bodyBytes))
	type Username struct {
		Name string `json:"name"`
	}
	var username Username
	err = json.Unmarshal(bodyBytes, &username)
	if err != nil {
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := db.QueryRow(context.Background(), "select role from users where username=$1", username.Name)
	err = result.Scan(&storedCreds.Role)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("User not exist!")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
