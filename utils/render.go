package utils

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, template_data map[string]string, template_name string) {
	if template_data == nil {
		template_data = make(map[string]string)
	}
	w.Header().Set("Content-Type", "text/html")
	template_text := Get_template(template_name)
	template := template.Must(template.New("data").Parse(template_text))
	err := template.Execute(w, template_data)
	if err != nil {
		panic(err)
	}
	// http.ServeFile(w, r, template.Execute(w, template_data))
}
