package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates insere os templates html na variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	// templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecutarTempleta renderiza uma página HTML na tela
func ExecutarTemplate(writer http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(writer, template, dados)
}