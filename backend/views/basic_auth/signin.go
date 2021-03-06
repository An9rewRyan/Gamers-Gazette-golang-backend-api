package basic_auth

import (
	"context"
	"d/go/structs"
	"d/go/utils/database"
	"d/go/utils/session"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyBytes))
	db, err := database.Connect_db()
	if err != nil {
		fmt.Println("err")
		return
	}
	defer db.Close()
	var creds structs.Credentials
	storedCreds := &structs.Credentials{}
	// Get the JSON body and decode into credentials
	err = json.Unmarshal(bodyBytes, &creds)
	if err != nil {
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from database
	result := db.QueryRow(context.Background(), "select password, role from accounts where username=$1", creds.Username)
	err = result.Scan(&storedCreds.Password, &storedCreds.Role)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == pgx.ErrNoRows {
			fmt.Println("User not exist!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("PAssword dont match!!")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(30 * time.Minute)

	// Set the token in the session map, along with the session information
	session.Sessions[sessionToken] = session.Session{
		Username: creds.Username,
		Expiry:   expiresAt,
		Role:     storedCreds.Role,
		Email:    storedCreds.Email,
		Bdate:    storedCreds.Bdate,
	}
	fmt.Println(session.Sessions)

	cookie := http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	}
	fmt.Println(cookie)

	b, err := json.Marshal(&cookie)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(b))
}
