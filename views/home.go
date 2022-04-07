package views

import (
	"d/go/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, nil, "index.html")
}
