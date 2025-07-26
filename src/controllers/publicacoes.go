package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/paulo-fabiano/apiDevBook/src/autenticacao"
	"github.com/paulo-fabiano/apiDevBook/src/banco"
	"github.com/paulo-fabiano/apiDevBook/src/modelos"
	"github.com/paulo-fabiano/apiDevBook/src/repositorios"
	"github.com/paulo-fabiano/apiDevBook/src/respostas"
)

func CriarPublicacao(writer http.ResponseWriter, request *http.Request) {

	usuarioID, err := autenticacao.ExtrairUsuarioID(request)
	if err != nil {
		respostas.Erro(writer, http.StatusUnauthorized, err)
		return
	}

	corpoRequest, err := ioutil.ReadAll(request.Body)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return
	}

	var publicacao modelos.Publicacao
	err = json.Unmarshal(corpoRequest, &publicacao)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return
	}

	publicacao.AutorID = usuarioID

	if err := publicacao.Preparar(); err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositoriosDePublicacoes(db)
	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}

	respostas.JSON(writer, http.StatusCreated, publicacao)

}

func BuscarPublicacao(writer http.ResponseWriter, request *http.Request) {
	
	parametros := mux.Vars(request)
	publicacaoID, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositoriosDePublicacoes(db)
	publicacao, err := repositorio.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return		
	}

	respostas.JSON(writer, http.StatusOK, publicacao)
}

func BuscarPublicacoes(writer http.ResponseWriter, request *http.Request) {
	
}

func AtualizarPublicacao(writer http.ResponseWriter, request *http.Request) {
	
}

func DeletarPublicacao(writer http.ResponseWriter, request *http.Request) {
	
}
