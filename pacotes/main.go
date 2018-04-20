package main

import (
    "fmt"

    "github.com/Zullus/cusrodego/pacotes/prefixo"
    "github.com/Zullus/cusrodego/pacotes/operadora"
)

//NomeDoUsuario é o nome do usuario do sistema
var NomeDoUsuario = "Edson"

func main() {

    fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
    fmt.Printf("Prefixo da Capital: %d\n\r", prefixo.Capital)
    fmt.Printf("Nome da Operadora: %s\n\r", operadora.NomeOperadora)
}