package api

import (
	"d/go/utils/database"
	"d/go/utils/html"
	"encoding/json"
	"fmt"
	"net/http"
)

func ApiArticles(w http.ResponseWriter, r *http.Request) {
	articles := database.Select_all_articles()
	b, err := json.Marshal(&articles)
	if err != nil {
		fmt.Println(err)
	}
	data := map[string]string{
		"Articles": string(b),
	}
	html.Render(w, r, data, "api.html")
}
