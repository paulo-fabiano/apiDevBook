package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/paulo-fabiano/apiDevBook/src/seguranca"
)

type Usuario struct {
	ID 			uint64		`json:"id,omitempty"`
	Nome 		string 		`json:"nome,omitempty"`
	Nick 		string 		`json:"nick,omitempty"`
	Email 		string 		`json:"email,omitempty"`
	Senha 		string 		`json:"senha,omitempty"`
	CriandoEm 	time.Time   `json:"criado_em,omitempty"`
}

func (usuario *Usuario) PrepararUsuario(etapa string) error {

	err := usuario.validar(etapa)
	if err != nil {
		return err
	}

	if err := usuario.formatar(etapa); err != nil {
		return err
	}

	return nil

}

func (usuario *Usuario) validar(etapa string) error {

	if usuario.Nome == "" {
		return errors.New("Campo Nome é obrigatório e está vazio")
	}

	if usuario.Nick == "" {
		return errors.New("Campo Nick é obrigátório e está vazio")
	}

	if usuario.Email == "" {
		return errors.New("Campo Email é obrigátório e está vazio")
	}

	err := checkmail.ValidateFormat(usuario.Email)
	if err != nil {
		return errors.New("O email inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("Campo Senha é obrigatória e está em branco")
	}

	return nil

}

func (usuario *Usuario) formatar(etapa string) error {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, err := seguranca.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
	
}