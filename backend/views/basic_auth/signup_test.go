package basic_auth

import (
	"bytes"
	"context"
	"d/go/errors"
	"d/go/structs"
	"d/go/utils/database"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup_test(w http.ResponseWriter, r *http.Request) {
	storedCreds := &structs.Credentials{}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	creds := structs.Credentials{}
	err = json.Unmarshal(bodyBytes, &creds)
	if err != nil {
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db, err := database.Connect_db()
	if err != nil {
		fmt.Println("err")
		return
	}
	defer db.Close()
	result := db.QueryRow(context.Background(), "select email, role from accounts where username=$1", creds.Username)
	err = result.Scan(&storedCreds.Email, &storedCreds.Role)
	if err == nil { //it means that user already exists and we need to tell frontend about it
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != pgx.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} //we continue only if user this this nickname does not exist
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err = db.Query(context.Background(), "insert into accounts values ($1, $2, 'user', $3, $4)", creds.Username, string(hashedPassword), creds.Email, creds.Bdate); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Failed to add user, ")
		fmt.Println(err)
		return
	} else {
		fmt.Println("Sucessfully added user!")
	}

	signin_link := "https://api-gamersgazette.herokuapp.com/auth/signin"
	req, _ := http.NewRequest("POST", signin_link, bytes.NewBuffer(bodyBytes))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending post:", err)
		errors.RespErr(w, err)
		return
	} else {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(bodyBytes), "Ola, sent response!")
		fmt.Fprint(w, string(bodyBytes))
		// fmt.Println("Sucessfully signed up!")
	}
}
