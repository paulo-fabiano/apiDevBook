package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
	ID uint64 `json:"id,omitempyt"`
	Titulo string `json:"titulo,omitempyt"`
	Conteudo string `conteudo,omitempty`
	AutorID uint64 `json:"autorId,omitempty"`
	AutorNick string `json:"autorNick,omitempty"`
	Curtidas uint64 `json:"curtidas"`
	CriandoEm time.Time `json:"criadoEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {

	if err := publicacao.validar(); err != nil {
		return err
	}

	publicacao.formatar()
	return nil

}

func (publicacao *Publicacao) validar() error {

	if publicacao.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("O contéudo é obrigatório e não pode estar em branco")	
	}

	return nil
}

func (publicacao *Publicacao) formatar() {

	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
	
}