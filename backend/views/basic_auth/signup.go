package basic_auth

import (
	"context"
	"d/go/structs"
	"d/go/utils/database"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	creds := &structs.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
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
	if _, err = db.Query(context.Background(), "insert into users values ($1, $2, 'admin', $3, $4)", creds.Username, string(hashedPassword), creds.Email, creds.Bdate); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Failed to add user, ")
		fmt.Println(err)
		return
	}
	fmt.Println("Sucessfully signed up!"))
	fmt.Fprintf(w, "Sucessfully signed up!")
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
}
