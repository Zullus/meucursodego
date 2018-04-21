package main

import (
	"fmt"
	"net/http"

	"github.com/Zullus/cursodego/banco_mongo/manipulador"
	"github.com/Zullus/cursodego/banco_mongo/repo"
)

func init() {

	fmt.Println("Iniciando o servidor")

}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()

	if err != nil {

		fmt.Println("Parando o servidor. Erro ao abrir o banco de dados: ", err.Error())

		return

	}

	err = repo.AbreSessaoComMongo()

	if err != nil {

		fmt.Println("Parando o servidor. Erro ao abrir o banco de dados MongoDB: ", err.Error())

		return

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Olá Mundo@")

	})

	http.HandleFunc("/funcao", manipulador.Funcao)

	http.HandleFunc("/ola", manipulador.Ola)

	http.HandleFunc("/local/", manipulador.Local)

	fmt.Println("Estou online...")

	http.ListenAndServe(":8181", nil)
}
