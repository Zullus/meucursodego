package repo

import (
	"github.com/jmoiron/sqlx"

	/*
		"github.com/go-sql-driver/mysql" Função iniciada antes e só para essa parte da aplicação
	*/
	_ "github.com/go-sql-driver/mysql"
)

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL a conexão com o banco de dados
func AbreConexaoComBancoDeDadosSQL() (err error) {

	err = nil

	Db, err = sqlx.Open("mysql", "cursogo:nx74205@tcp(192.168.0.10:3306)/cursogo")

	if err != nil {
		return
	}

	err = Db.Ping() //Faz a chamada com o banco de dados

	return
}
