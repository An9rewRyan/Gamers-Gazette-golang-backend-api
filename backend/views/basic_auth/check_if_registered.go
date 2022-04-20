package basic_auth

import (
	"context"
	"d/go/structs"
	"d/go/utils/database"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
)

func Check_if_registered(w http.ResponseWriter, r *http.Request) {
	storedCreds := &structs.Credentials{}
	creds := structs.Credentials{}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(bodyBytes, &creds)
	if err != nil {
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
	}

	db, err := database.Connect_db()
	if err != nil {
		fmt.Println("err")
		return
	}
	defer db.Close()

	result := db.QueryRow(context.Background(), "select role from users where username=$1, email=$2", creds.Username, creds.Email)
	err = result.Scan(&storedCreds.Email, &storedCreds.Role)
	if err == nil { //already registered
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != pgx.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError) //iternal error
		return
	} else {
		w.WriteHeader(http.StatusOK) //not registered
		return
	}
}
