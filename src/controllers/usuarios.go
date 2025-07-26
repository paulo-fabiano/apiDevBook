package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/paulo-fabiano/apiDevBook/src/autenticacao"
	"github.com/paulo-fabiano/apiDevBook/src/banco"
	"github.com/paulo-fabiano/apiDevBook/src/modelos"
	"github.com/paulo-fabiano/apiDevBook/src/repositorios"
	"github.com/paulo-fabiano/apiDevBook/src/respostas"
	"github.com/paulo-fabiano/apiDevBook/src/seguranca"
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
		Data interface{} `json:"mensagem"`
	}{
		Data: usuario,
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

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(request)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return	
	}

	// Se o ID do usuario for diferente do ID do token então o usuário não consegue fazer uma atualização dos seus dados
	// então, um usuário não pode alterar as informações do outro
	if usuarioID != usuarioIDNoToken {
		respostas.Erro(writer, http.StatusForbidden, errors.New("Você não pode atualizar outros usuários"))
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

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(request)
	if err != nil {
		respostas.Erro(writer, http.StatusUnauthorized, err)
		return	
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(writer, http.StatusForbidden, errors.New("Você não pode excluir outros usuários"))
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

func SeguirUsuario(writer http.ResponseWriter, request *http.Request) {
	
	seguidorID, err := autenticacao.ExtrairUsuarioID(request)
	if err != nil {
		respostas.Erro(writer, http.StatusUnauthorized, err)
		return	
	}

	paramatros := mux.Vars(request)
	usuarioID, err := strconv.ParseUint(paramatros["id"], 10, 64)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return		
	}

	if seguidorID == usuarioID {
		respostas.Erro(writer, http.StatusForbidden, errors.New("Não é possível seguir você mesmo"))
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	err = repositorio.Seguir(usuarioID, seguidorID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}

	respostas.JSON(writer, http.StatusNoContent, nil)

}

func PararDeSeguirUsuario(writer http.ResponseWriter, request *http.Request) {
	
	seguidorID, err := autenticacao.ExtrairUsuarioID(request)
	if err != nil {
		respostas.Erro(writer, http.StatusUnauthorized, err)
		return	
	}

	paramatros := mux.Vars(request)
	usuarioID, err := strconv.ParseUint(paramatros["id"], 10, 64)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return		
	}

	if seguidorID == usuarioID {
		respostas.Erro(writer, http.StatusForbidden, errors.New("Não é possível para de seguir você mesmo"))
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	err = repositorio.PararDeSeguir(usuarioID, seguidorID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}

	respostas.JSON(writer, http.StatusNoContent, nil)

}

func BuscarSeguidores(writer http.ResponseWriter, request *http.Request) {
	

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
	seguidores, err := repositorio.BuscarSeguidores(usuarioID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}

	respostas.JSON(writer, http.StatusOK, seguidores)

}

func BuscarSeguindo(writer http.ResponseWriter, request *http.Request) {
	

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
	seguindo, err := repositorio.BuscarSeguindo(usuarioID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}

	respostas.JSON(writer, http.StatusOK, seguindo)

}

func AtualizarSenhaUsuario(writer http.ResponseWriter, request *http.Request) {
	
	parametros := mux.Vars(request)
	usuarioID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return		
	}

	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(request)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return	
	}

	if usuarioIDNoToken != usuarioID {
		respostas.Erro(writer, http.StatusForbidden, errors.New("Não é possível atualizar a de um usuário que não é o seu!"))
	}

	corpoRequest, err := ioutil.ReadAll(request.Body)

	var senha modelos.Senha
	err = json.Unmarshal(corpoRequest, &senha)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
	}


	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	senhaSalvaNoBanco, err := repositorio.BuscarSenha(usuarioID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}

	if err := seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); err != nil {
		respostas.Erro(writer, http.StatusUnauthorized, errors.New("A senha atual está incorreta"))
		return	
	}

	senhaComHash, err := seguranca.Hash(senha.Nova)
	if err != nil {
		respostas.Erro(writer, http.StatusUnauthorized, err)
		return	
	}
	
	if err := repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return		
	}

	respostas.JSON(writer, http.StatusNoContent, nil)

}