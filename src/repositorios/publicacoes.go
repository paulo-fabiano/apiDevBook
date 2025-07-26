package repositorios

import (
	"database/sql"

	"github.com/paulo-fabiano/apiDevBook/src/modelos"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositoriosDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{
		db: db,
	}
}

func (repo Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {

	stmt, err := repo.db.Prepare(`
		insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.Conteudo)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil

}

func (repo Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {

	linha, err := repo.db.Query(`
		select p.*, u.nick from
		publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = ?
	`, publicacaoID)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer linha.Close()

	var publicacao modelos.Publicacao
	for linha.Next() {
		err := linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriandoEm,
			&publicacao.AutorNick,
		)
		if err != nil {
			return modelos.Publicacao{}, err
		}
	}

	return publicacao, nil

}

func (repo Publicacoes) Buscar(usurioID uint64) ([]modelos.Publicacao, error) {

	linhas, err := repo.db.Query(`
		select distinct p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.autor_id
		inner join seguidores s on p.autor_id = s.usuario_id
		where u.id = ? or s.seguidor_id = ?
		order by 1 desc
	`, usurioID, usurioID)
	if err != nil {
		return []modelos.Publicacao{}, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		err := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriandoEm,
			&publicacao.AutorNick,
		)
		if err != nil {
			return []modelos.Publicacao{}, err
		}

		publicacoes = append(publicacoes, publicacao)

	}

	return publicacoes, nil

}

func (repo Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {

	stmt, err := repo.db.Prepare(`
		update publicacoes set titulo = ?, conteudo = ?, where id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}

	return nil

}

// Deletar deleta uma publicação do banco de dados
func (repo Publicacoes) Deletar(publicacaoID uint64) error {

	stmt, err := repo.db.Prepare(`
		delete from publicacoes where id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(publicacaoID); err != nil {
		return err
	}

	return nil

}