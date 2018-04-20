package main

import (
	"fmt"
	"net/http"

	"github.com/Zullus/cusrodego/banco_sql/manipulador"
	"github.com/Zullus/cusrodego/banco_sql/repo"
)

func int() { //Executa antes do main

	fmt.Println("Carregando o servidor")
}

func main() {

	err := repo.AbreConexaoComBandoDeDadosSQL()

	if err != nil {

		fmt.Println("[main] Servidor parado, erro ao abrir o banco de dados: ", err.Error())

		return

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Olá Mundo@")

	})

	http.HandleFunc("/funcao", manipulador.Funcao)

	http.HandleFunc("/ola", manipulador.Ola)

	http.HandleFunc("/local", manipulador.Local)

	fmt.Println("Estou online...")

	http.ListenAndServe(":8181", nil)
}
