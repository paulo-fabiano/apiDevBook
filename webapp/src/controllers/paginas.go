package controllers

import (
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/src/utils"
)

func CarregarTelaDeLogin(writer http.ResponseWriter, request *http.Request) {

	utils.ExecutarTemplate(writer, "login.html", nil) // O terceiro parâmetro é nil por quê não iremos inserir nenhum dado variável na tela

}

func CarregarPaginaDeCadastroDeUsuarios(writer http.ResponseWriter, request *http.Request) {
	utils.ExecutarTemplate(writer, "cadastro.html", nil)
}