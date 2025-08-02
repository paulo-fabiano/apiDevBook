package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/src/autenticacao"
	"github.com/paulo-fabiano/apiDevBook/src/banco"
	"github.com/paulo-fabiano/apiDevBook/src/modelos"
	"github.com/paulo-fabiano/apiDevBook/src/repositorios"
	"github.com/paulo-fabiano/apiDevBook/src/respostas"
	"github.com/paulo-fabiano/apiDevBook/src/seguranca"
)

func Login(writer http.ResponseWriter, request *http.Request) {

	corpoRequest, err := ioutil.ReadAll(request.Body)
	if err != nil {
		respostas.Erro(writer, http.StatusBadRequest, err)
		return
	}

	var usuario modelos.Usuario
	if err := json.Unmarshal(corpoRequest, &usuario); err != nil {
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
	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return		
	}

	// Validando login
	err = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha)
	if err != nil {
		respostas.Erro(writer, http.StatusUnauthorized, err)
		return		
	}

	token, err := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		respostas.Erro(writer, http.StatusInternalServerError, err)
		return	
	}
	writer.Write([]byte(token))

}