package handlers

import(
	"text/template"
)

const (
	indexTemplatePath = "./templates/index.html"
	errorTemplatePath = "./templates/error.html"
)

var templates Templates

func LoadTemplates() {
	templates.Index = template.Must(template.ParseFiles(indexTemplatePath))
	templates.Error = template.Must(template.ParseFiles(errorTemplatePath))
}