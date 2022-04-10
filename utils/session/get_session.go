package session

import (
	"net/http"
)

func Get_session(w http.ResponseWriter, r *http.Request) (Session, string) {
	// We can obtain the session token from the requests cookies, which come with every request
	var s Session
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return s, "unauthorized"
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return s, "bad request"
	}
	sessionToken := c.Value

	// We then get the session from our session map
	userSession, exists := Sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return s, "unauthorized"
	}
	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if userSession.IsExpired() {
		delete(Sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return s, "unauthorized"
	}

	// If the session is valid, return the welcome message to the user
	return userSession, "valid"
}
