package modelos

import "time"

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
	ID uint64 `json:"id,omitempyt"`
	Titulo string `json:"titulo,omitempyt"`
	Conteudo string `conteudo,omitempty`
	AutorID uint64 `json:"autorId,omitempty"`
	AutorNick uint64 `json:"autorNick,omitempty"`
	Curtidas uint64 `json:"curtidas"`
	CriandoEm time.Time `json:"criadoEm,omitempty"`
}