package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/paulo-fabiano/apiDevBook/src/config"
)

func Conectar() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}