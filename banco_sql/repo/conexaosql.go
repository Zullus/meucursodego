package repo

import(

    "github.com/jmoiron/sqlx"

    /*
    github.com/go-driver-mysql usado diretamente pela aplicação
    */
    _ "github.com/go-driver-mysql" //
)

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBandoDeDadosSQL conexao com o banco de dados
func AbreConexaoComBandoDeDadosSQL() (err error) {

    err = nil //forçando o erro a ser nulo
    Db, err = sqlx.Open("mysql", "curdogo@tcp(192.168.0.10:3306)/curdogo")

    if err != nil{

        return

    }

    err =Db.Ping()

    if err != nil{

        return

    }

    return
}