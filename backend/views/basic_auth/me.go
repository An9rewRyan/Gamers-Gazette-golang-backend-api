package basic_auth

import (
	"d/go/utils/session"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bodyBytes))
	type Cookie struct {
		Session_token string `json:"session_token"`
	}
	var cookie Cookie
	err = json.Unmarshal(bodyBytes, &cookie)
	if err != nil {
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := cookie.Session_token

	// We then get the session from our session map
	userSession, exists := session.Sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if userSession.IsExpired() {
		delete(session.Sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If the session is valid, return the welcome message to the user
	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.Username)))
}
