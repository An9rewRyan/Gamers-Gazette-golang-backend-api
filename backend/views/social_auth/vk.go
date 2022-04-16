package social_auth

import (
	"bytes"
	"crypto/tls"
	"d/go/errors"
	"d/go/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

var redirectURI = "https://gamersgazette.herokuapp.com/signup/vk"
var clientID = "8134856"
var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var client = &http.Client{Transport: tr}
var scope = []string{"account", "email", "bdate"}
var scopeTemp = strings.Join(scope, "+")
var clientSecret = "7Vw4ALUIHMLPpHTKiRlG"

func Vk_redir(w http.ResponseWriter, r *http.Request) {
	newUrl := fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s", clientID, redirectURI, scopeTemp)
	fmt.Println(newUrl)
	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}

func Vk_get_data(w http.ResponseWriter, r *http.Request) {
	resp_bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	type url_response struct {
		Url_with_code string `json:"url_with_code"`
	}
	resp_url := url_response{}
	err = json.Unmarshal(resp_bytes, &resp_url)
	if err != nil {
		fmt.Println("Decode error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(resp_url.Url_with_code)
	url_resp, err := url.ParseRequestURI(resp_url.Url_with_code)
	if err != nil {
		fmt.Println(err)
		return
	}

	code := url_resp.Query().Get("code")
	if code == "" {
		errors.RespErr(w, fmt.Errorf("code query param is not provided"))
		return
	}
	fmt.Println(code, r.URL.RequestURI())
	url := fmt.Sprintf("https://oauth.vk.com/access_token?grant_type=authorization_code&code=%s&redirect_uri=%s&client_id=%s&client_secret=%s", code, redirectURI, clientID, clientSecret)
	req, _ := http.NewRequest("POST", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		errors.RespErr(w, err)
		return
	}
	defer resp.Body.Close()
	resp_bytes, _ = ioutil.ReadAll(resp.Body)
	email := gjson.Get(string(resp_bytes), "email")
	fmt.Println(string(resp_bytes))
	url = fmt.Sprintf("https://api.vk.com/method/users.get?access_token=%s&fields=bdate&user_id=%s&v=5.131", gjson.Get(string(resp_bytes), "access_token"), gjson.Get(string(resp_bytes), "user_id"))
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		errors.RespErr(w, err)
		return
	}
	resp, err = client.Do(req)
	if err != nil {
		errors.RespErr(w, err)
		return
	}
	defer resp.Body.Close()
	resp_bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.RespErr(w, err)
		return
	}
	fmt.Println(string(resp_bytes))
	user := structs.Soc_auth_data{
		Username:  fmt.Sprintf("%s %s", gjson.Get(string(resp_bytes), "response.#.first_name").String(), gjson.Get(string(resp_bytes), "response.#.last_name").String()),
		BirthDate: gjson.Get(string(resp_bytes), "response.#.bdate").String(),
		Email:     email.String(),
	}
	b, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	signin_link := "https://gamersgazette.herokuapp.com/signup/vk"
	req, _ = http.NewRequest("POST", signin_link, bytes.NewBuffer(b))
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("Error while sending post:", err)
		errors.RespErr(w, err)
		return
	}
}
