package social_auth

import (
	"crypto/tls"
	"d/go/structs"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func Vk(w http.ResponseWriter, r *http.Request) {
	redirectURI := "https://gamersgazette.herokuapp.com/auth/me"
	clientID := "8134856"
	clientSecret := "7Vw4ALUIHMLPpHTKiRlG"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	scope := []string{"account", "email", "bdate"}
	scopeTemp := strings.Join(scope, "+")
	url := fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s", clientID, redirectURI, scopeTemp)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		respErr(w, err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		respErr(w, err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		respErr(w, err)
		return
	}
	fmt.Fprint(w, string(bytes), r.URL.Query().Get("code"))
	code := r.URL.Query().Get("code")
	if code == "" {
		respErr(w, fmt.Errorf("code query param is not provided"))
		return
	}
	fmt.Println(code, r.URL.RequestURI())
	url = fmt.Sprintf("https://oauth.vk.com/access_token?grant_type=authorization_code&code=%s&redirect_uri=%s&client_id=%s&client_secret=%s", code, redirectURI, clientID, clientSecret)
	req, _ = http.NewRequest("POST", url, nil)
	resp, err = client.Do(req)
	if err != nil {
		respErr(w, err)
		return
	}
	defer resp.Body.Close()
	bytes, _ = ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(bytes))
	email := gjson.Get(string(bytes), "email")
	url = fmt.Sprintf("https://api.vk.com/method/users.get?access_token=%s&fields=bdate&user_id=%s&v=5.131", gjson.Get(string(bytes), "access_token"), gjson.Get(string(bytes), "user_id"))
	fmt.Println(url)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		respErr(w, err)
		return
	}
	resp, err = client.Do(req)
	if err != nil {
		respErr(w, err)
		return
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		respErr(w, err)
		return
	}
	fmt.Println(gjson.Get(string(bytes), "response.#.first_name"), gjson.Get(string(bytes), "response.#.last_name"), gjson.Get(string(bytes), "response.#.bdate"), email)
	fmt.Fprint(w, gjson.Get(string(bytes), "response.#.first_name"), gjson.Get(string(bytes), "response.#.last_name"), gjson.Get(string(bytes), "response.#.bdate"), email)

	user := structs.Soc_auth_data{
		Username:  gjson.Get(string(bytes), "response.#.first_name").String() + " " + gjson.Get(string(bytes), "response.#.last_name").String(),
		BirthDate: gjson.Get(string(bytes), "response.#.bdate").String(),
		Email:     email.String(),
	}
	b, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(b))
}

// func Me(w http.ResponseWriter, r *http.Request) {
// 	redirectURI := "https://gamersgazette.herokuapp.com/auth/me"
// 	clientID := "8134856"
// 	clientSecret := "7Vw4ALUIHMLPpHTKiRlG"
// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
// 	client := &http.Client{Transport: tr}
// 	code := r.URL.Query().Get("code")
// 	if code == "" {
// 		respErr(w, fmt.Errorf("code query param is not provided"))
// 		return
// 	}
// 	fmt.Println(code, r.URL.RequestURI())
// 	url := fmt.Sprintf("https://oauth.vk.com/access_token?grant_type=authorization_code&code=%s&redirect_uri=%s&client_id=%s&client_secret=%s", code, redirectURI, clientID, clientSecret)
// 	req, _ := http.NewRequest("POST", url, nil)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		respErr(w, err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	bytes, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Fprint(w, string(bytes))
// 	email := gjson.Get(string(bytes), "email")
// 	url = fmt.Sprintf("https://api.vk.com/method/users.get?access_token=%s&fields=bdate&user_id=%s&v=5.131", gjson.Get(string(bytes), "access_token"), gjson.Get(string(bytes), "user_id"))
// 	fmt.Println(url)
// 	req, err = http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		respErr(w, err)
// 		return
// 	}
// 	resp, err = client.Do(req)
// 	if err != nil {
// 		respErr(w, err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	bytes, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		respErr(w, err)
// 		return
// 	}
// 	fmt.Println(gjson.Get(string(bytes), "response.#.first_name"), gjson.Get(string(bytes), "response.#.last_name"), gjson.Get(string(bytes), "response.#.bdate"), email)
// 	fmt.Fprint(w, gjson.Get(string(bytes), "response.#.first_name"), gjson.Get(string(bytes), "response.#.last_name"), gjson.Get(string(bytes), "response.#.bdate"), email)
// }

func respErr(w http.ResponseWriter, err error) {
	_, er := io.WriteString(w, err.Error())
	if er != nil {
		log.Println(err)
	}
}