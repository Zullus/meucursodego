package main

import "fmt"

var endereco string //Assume endereco = ""
var telefone = "9999-9999"
var quantidade int //alias para int32 em máquina 32bits e int64 em máquinas 64bits assume quantidade = 0
var comprou bool //assume comprou = false
var valor float64 //assume valor = 0.00
var palavras rune //para unicode

func main() {

    fmt.Printf("endereco: %s\r\n", endereco)

    fmt.Printf("quantidade: %d\r\n", quantidade)

    fmt.Printf("comprou: %v\r\n", comprou)

    //fmt.Printf("valor: %\r\n")

    //fmt.Printf("palavras: %\n\r")
}