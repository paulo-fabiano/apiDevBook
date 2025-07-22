package rotas

import (
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/src/controllers"
)

var rotaLogin = Rota{
	URI: "/login",
	Metodo: http.MethodPost,
	Funcao: controllers.Login,
	RequerAutenticacao: false,
}