package api

import (
	"d/go/utils/database"
	"d/go/utils/html"
	"d/go/utils/session"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiArticleDelete(w http.ResponseWriter, r *http.Request) {
	session, status := session.Get_session(w, r)
	if status == "unauthorized" {
		fmt.Fprint(w, "You need to sign in to perform this action!")
		return
	}
	if status == "bad request" {
		fmt.Fprint(w, "Error on backend! (Iternal server error)")
		return
	}
	if status == "valid" && session.Role == "admin" {
		vars := mux.Vars(r)
		article, err := database.Select_article(vars["id"])
		if err != nil {
			data := map[string]string{
				"Articles": "404! No article with such id is found!",
			}
			html.Render(w, r, data, "api.html")
			return
		}
		b, err := json.Marshal(&article)
		if err != nil {
			fmt.Println(err)
		}
		_, err = database.Delete_article(vars["id"])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, string(b))
	} else {
		fmt.Fprint(w, "You are not allowed, to acess this source! If you are admin, please sign in your account!")
	}
}
