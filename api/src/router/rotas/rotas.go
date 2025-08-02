package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulo-fabiano/apiDevBook/src/middlewares"
)

type Rota struct {
	URI						string
	Metodo 					string
	Funcao 					func(http.ResponseWriter, *http.Request)
	RequerAutenticacao 		bool
}

func Configurar(r *mux.Router) *mux.Router {

	rotas := RotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...) // ... quer dizer que ele tem que dar um append para cada tipo do slice de rotasPublicacoes

	for _, rota := range rotas {

		if rota.RequerAutenticacao { // Se for true
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	return r
	
}