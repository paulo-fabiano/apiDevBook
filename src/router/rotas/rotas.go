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

	for _, rota := range rotas {

		if rota.RequerAutenticacao { // Se for true
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	return r
	
}