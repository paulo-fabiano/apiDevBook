package middlewares

import (
	"log"
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/src/autenticacao"
	"github.com/paulo-fabiano/apiDevBook/src/respostas"
)

// Loger escreve informações das requisições
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {

	return func (writer http.ResponseWriter, request *http.Request)  {
		log.Printf("\n %s %s %s", request.Method, request.RequestURI, request.Host)
		proximaFuncao(writer, request)
	}
}

// Autenticar verifica se a rota é pública ou privada
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		if err := autenticacao.ValidarToken(request); err != nil {
			respostas.Erro(writer, http.StatusUnauthorized, err)
			return
		}
		proximaFuncao(writer, request)
	}

}