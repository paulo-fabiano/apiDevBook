package controllers

import (
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/utils"
)

func CarregarTelaDeLogin(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("Aqui ta normal"))
	utils.ExecutarTemplate(writer, "login.html", nil) // O terceiro parâmetro é nil por quê não iremos inserir nenhum dado variável na tela

}