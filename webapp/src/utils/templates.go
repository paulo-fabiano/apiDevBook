package utils

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

// CarregarTemplates insere os templates html na variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// // ExecutarTempleta renderiza uma página HTML na tela
// func ExecutarTemplate(writer http.ResponseWriter, template string, dados interface{}) {
// 	templates.ExecuteTemplate(writer, template, dados)
// }
// ExecutarTemplate renderiza uma página HTML na tela
func ExecutarTemplate(writer http.ResponseWriter, templateName string, dados interface{}) {
    if templates == nil {
        http.Error(writer, "Template não carregado.", http.StatusInternalServerError)
        return
    }

    err := templates.ExecuteTemplate(writer, templateName, dados)
    if err != nil {
        log.Printf("Erro ao executar o template '%s': %v", templateName, err)
        http.Error(writer, "Erro ao renderizar a página.", http.StatusInternalServerError)
    }
}