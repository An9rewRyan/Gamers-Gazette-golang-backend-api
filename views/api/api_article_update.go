package api

import (
	"d/go/structs"
	"d/go/utils/database"
	"d/go/utils/html"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiArticleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	article := structs.Article_create{
		Title:      r.URL.Query().Get("title"),
		Pub_date:   r.URL.Query().Get("pub_date"),
		Image_url:  r.URL.Query().Get("image_url"),
		Content:    r.URL.Query().Get("content"),
		Src_link:   r.URL.Query().Get("src_link"),
		Site_alias: r.URL.Query().Get("site_alias"),
	}
	err := database.Update_article(vars["id"], article)
	if err != nil {
		data := map[string]string{
			"Articles": "404! No article with such id is found!",
		}
		html.Render(w, r, data, "api.html")
		return
	}
	article_new, err := database.Select_article(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(&article_new)
	if err != nil {
		fmt.Println(err)
	}
	data := map[string]string{
		"Articles": "Updated artcle: " + string(b),
	}
	html.Render(w, r, data, "api.html")
}
