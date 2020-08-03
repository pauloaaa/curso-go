package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// ConectaComBancoDeDados Conecta na base de dados
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=paulo_arruda_loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
