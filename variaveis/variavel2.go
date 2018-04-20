package main

import "fmt"

//Variável disponíveis para todo o programa
var (
    //Endereco é importante e tem de ser público
    Endereco    string
    //Telefone é imorante e tem de ser pública
    Telefone    string
    quantidade  int
    comprou     bool
    valor       float64
    palavras    rune
)

func main() {

    teste := "Valor de teste" //Variavel estrita a funcção
    fmt.Printf("endereco: %s\r\n", Endereco)
    fmt.Printf("quantidade: %d\r\n", quantidade)
    fmt.Printf("comprou: %v\r\n", comprou)
    fmt.Printf("valor: %v\r\n", valor)
    fmt.Printf("palavras: %v\n\r", palavras)
    fmt.Printf("O valor do teste é: %v\r\n", teste)
}