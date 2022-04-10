package api

import (
	"d/go/structs"
	"d/go/utils/database"
	"d/go/utils/session"
	"encoding/json"
	"fmt"
	"net/http"
)

func ApiArticleCreate(w http.ResponseWriter, r *http.Request) {
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
		article := structs.Article_create{
			Title:      r.URL.Query().Get("title"),
			Pub_date:   r.URL.Query().Get("pub_date"),
			Image_url:  r.URL.Query().Get("image_url"),
			Content:    r.URL.Query().Get("content"),
			Src_link:   r.URL.Query().Get("src_link"),
			Site_alias: r.URL.Query().Get("site_alias"),
		}

		database.Write_article_to_db(article)
		b, err := json.Marshal(&article)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, string(b))
	} else {
		fmt.Fprint(w, "You are not allowed, to acess this source! If you are admin, please sign in your account!")
	}
}

// example query for testing
// http://127.0.0.1:8000/api/articles/create?title=title1&content=content1&pub_date=2021-01-01%2017:01:17&site_alias=test&image_url=test/images&src_link=test.com/article1
