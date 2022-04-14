package basic_auth

import (
	"bytes"
	"context"
	"crypto/tls"
	"d/go/errors"
	"d/go/structs"
	"d/go/utils/database"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var client = &http.Client{Transport: tr}

func Signup(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyBytes))
	creds := &structs.Credentials{}
	err = json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(creds)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := database.Connect_db()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Next, insert the username, along with the hashed password into the database
	if _, err = db.Query(context.Background(), "insert into users values ($1, $2, 'user', $3, $4)", creds.Username, string(hashedPassword), creds.Email, creds.Bdate); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Failed to add user, ")
		fmt.Println(err)
		return
	} else {
		fmt.Println("Sucessfully added user!")
	}

	signin_link := "https://api-gamersgazette.herokuapp.com/auth/signin"
	req, _ := http.NewRequest("POST", signin_link, bytes.NewBuffer(bodyBytes))
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("Error while sending post:", err)
		errors.RespErr(w, err)
		return
	} else {
		fmt.Println("Sucessfully signed up!")
		fmt.Fprint(w, "Sucessfully signed up!")
	}
}
