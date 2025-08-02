package router

import (
	"github.com/gorilla/mux"
	"github.com/paulo-fabiano/apiDevBook/src/router/rotas"
)

// GerarRouter vai retornar um router com as rotas configuradas
func GerarRouter() *mux.Router {

	r := mux.NewRouter()
	return rotas.Configurar(r)
	
}