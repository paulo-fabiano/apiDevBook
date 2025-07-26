package repositorios

import (
	"database/sql"
	"fmt"

	"github.com/paulo-fabiano/apiDevBook/src/modelos"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar cria um novo usuário
func (repo Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	stmt, err := repo.db.Prepare("insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)")
	if err != nil {
		return 0, nil
	}
	defer stmt.Close()

	resultato, err := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultato.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar retornar todos os usuários que atendem a um nome ou nick específico
func (repo Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, err := repo.db.Query(
		"select id, nome, nick, email, criado_em from usuarios where nome LIKE ? or nick like ?",
		nomeOuNick, nomeOuNick,
	)
	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		err := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriandoEm,
		)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

// BuscarPorID busca um usuário no banco de dados
func (repo Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, err := repo.db.Query(
		"select id, nome, nick, email, criado_em from usuarios where id = ?",
		ID,
	)
	if err != nil {
		return modelos.Usuario{}, err
	}

	defer linhas.Close()

	var usuario modelos.Usuario
	for linhas.Next() {
		err := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriandoEm,
		)
		if err != nil {
			return modelos.Usuario{}, err
		}
	
	}

	return usuario, nil

}

// Atualizar atualiza um usuário do banco de dados
func (repo Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {

	stmt, err := repo.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID)
	if err != nil {
		return err
	}

	return nil

}

// Deletar apaga um usuário do banco de dados
func (repo Usuarios) Deletar(ID uint64) error {

	stmt, err := repo.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

// BuscarPorEmail faz uma busca por email e retorna o objeto, se houver
func (repo Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {

	linha, err := repo.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()

	var usuario modelos.Usuario

	for linha.Next() {
		err := linha.Scan(&usuario.ID, &usuario.Senha)
		if err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil

}

func (repo Usuarios) Seguir(usuarioID, seguidorID uint64) error {

	stmt, err := repo.db.Prepare("insert ignore into seguidores(usuario_id, seguidor_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (repo Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {

	stmt, err := repo.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (repo Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, err := repo.db.Query("select u.id, u.nome, u.nick, u.email, u.criado_em from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {

		var usuario modelos.Usuario
		if err := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.CriandoEm,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

func (repo Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, err := repo.db.Query("select u.id, u.nome, u.nick, u.email, u.criado_em from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {

		var usuario modelos.Usuario
		if err := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.CriandoEm,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
	
}

func (repo Usuarios) BuscarSenha(usuarioID uint64) (string, error) {

	linha, err := repo.db.Query(`
		select senha from usuarios where id = ?
	`)
	if err != nil {
		return "", err
	}
	defer linha.Close()

	var usuario modelos.Usuario

	for linha.Next() {
		err = linha.Scan(&usuario.Senha)
		if err != nil {
			return "", err
		}
	}

	return usuario.Senha, err

}

func (repo Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {

	stmt, err := repo.db.Prepare(`
		update usuarios set senha = ? where id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(usuarioID, senha); err != nil {
		return err
	}

	return nil

}