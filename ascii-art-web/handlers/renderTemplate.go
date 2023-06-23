package handlers

import (
	"net/http"
	"text/template"
)

type Templates struct {
	Index *template.Template
	Error *template.Template
}

func RenderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}