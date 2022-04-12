package basic_auth

import (
	"d/go/utils/session"
	"fmt"
	"net/http"
)

func LoggedIn(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

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
