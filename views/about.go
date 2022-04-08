package views

import (
	"d/go/utils/html"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"Title":   "World Cup",
		"Message": "FIFA will never regret it",
	}
	html.Render(w, r, data, "about.html")
}
