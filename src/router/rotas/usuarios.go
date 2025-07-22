package rotas

import (
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/src/controllers"
)

var RotasUsuarios = []Rota{
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{id}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{id}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{id}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}