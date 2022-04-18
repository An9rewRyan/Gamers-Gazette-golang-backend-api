package basic_auth

import (
	"d/go/utils/session"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bodyBytes))
	fmt.Println(session.Sessions)
	type Cookie struct {
		Session_token string `json:"session_cookie"`
	}
	var cookie Cookie
	err = json.Unmarshal(bodyBytes, &cookie)
	if err != nil {
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := cookie.Session_token
	delete(session.Sessions, sessionToken)
}
