package basic

import (
	"d/go/utils/html"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	html.Render(w, r, nil, "index.html")
}
