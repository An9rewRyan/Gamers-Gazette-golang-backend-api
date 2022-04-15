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
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	creds := structs.Credentials{}
	err = json.Unmarshal(bodyBytes, &creds)
	if err != nil {
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	if _, err = db.Query(context.Background(), "insert into users values ($1, $2, 'user', $3, $4)", creds.Username, string(hashedPassword), creds.Email, creds.Bdate); err != nil {
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
			log.Fatal(err)
		}
		fmt.Println(string(bodyBytes), "Ola!")
		fmt.Fprint(w, string(bodyBytes))
		fmt.Println("Sucessfully signed up!")
		fmt.Fprint(w, "Sucessfully signed up!")
	}
}
