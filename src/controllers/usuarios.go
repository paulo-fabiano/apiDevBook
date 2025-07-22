package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/paulo-fabiano/apiDevBook/src/banco"
	"github.com/paulo-fabiano/apiDevBook/src/modelos"
	"github.com/paulo-fabiano/apiDevBook/src/repositorios"
	"github.com/paulo-fabiano/apiDevBook/src/respostas"
)

func CriarUsuario(writer http.ResponseWriter, request *http.Request) {

	corpoRequest, err := ioutil.ReadAll(request.Body)
	if err != nil {
		respostas.Erro(writer, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario
	err = json.Unmarshal(corpoRequest, &usuario)
	if err != nil {
		respostas.Erro(writer, http.StatusBadGateway, err)
		return
	}

	// Fazendo as validações de usuário
	err = usuario.PrepararUsuario("cadastro")
	if err != nil {
		respostas.Erro(writer, http.StatusBadGateway, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuario.ID, err = repositorio.Criar(usuario)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(writer, http.StatusCreated, struct{
		Mensagem string `json:"mensagem"`
	}{
		Mensagem: fmt.Sprintf("ID %d inserido com sucesso", usuario.ID),
	})

}

func BuscarUsuarios(writer http.ResponseWriter, request *http.Request) {

	nomeOuNick := strings.ToLower(request.URL.Query().Get("usuario"))

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, err := repositorio.Buscar(nomeOuNick)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}
	
	respostas.JSON(writer, http.StatusOK, struct{
		Data interface{} `json:"data"`
	}{
		Data: usuarios,
	})
}

func BuscarUsuario(writer http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)
	usuarioID, err := strconv.ParseUint(parametros["id"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuario, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}
	
	respostas.JSON(writer, http.StatusOK, struct{
		Data interface{} `json:"data"`
	}{
		Data: usuario,
	})

}

func AtualizarUsuario(writer http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)
	usuarioID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return		
	}

	corpoRequest, err := ioutil.ReadAll(request.Body)
	if err != nil {
		respostas.Erro(writer, http.StatusUnprocessableEntity, err)
		return	
	}
	
	var usuario modelos.Usuario
	err = json.Unmarshal(corpoRequest, &usuario)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return				
	}

	err = usuario.PrepararUsuario("atualizarUsuario")
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

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	err = repositorio.Atualizar(usuarioID, usuario)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}

	respostas.JSON(writer, http.StatusNoContent, nil)

}

func DeletarUsuario(writer http.ResponseWriter, request *http.Request) {

	paramatros := mux.Vars(request)
	usuarioID, err := strconv.ParseUint(paramatros["id"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	err = repositorio.Deletar(usuarioID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}

	respostas.JSON(writer, http.StatusNoContent, nil)
	
}
