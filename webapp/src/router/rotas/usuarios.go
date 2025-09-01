package rotas

import (
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI: "/criar-usuario",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaDeCadastroDeUsuarios,
		RequerAutenticacao: false,
	},
}