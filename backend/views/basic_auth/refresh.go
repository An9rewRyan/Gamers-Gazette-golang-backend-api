package basic_auth

import (
	"d/go/utils/session"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code from this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	userSession, exists := session.Sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.IsExpired() {
		delete(session.Sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// If the previous session is valid, create a new session token for the current user
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(30 * time.Minute)

	// Set the token in the session map, along with the user whom it represents
	session.Sessions[newSessionToken] = session.Session{
		Username: userSession.Username,
		Expiry:   expiresAt,
		Role:     userSession.Role,
		Email:    userSession.Email,
	}

	// Delete the older session token
	delete(session.Sessions, sessionToken)

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(30 * time.Minute),
	})
}
