package main

import (
	"fmt"
)

func main() {

	var nuns []int

	fmt.Println(nuns, len(nuns), cap(nuns))

	nuns = make([]int, 5) //inicia um slice

	fmt.Println(nuns, len(nuns), cap(nuns))

	capitais := []string{"Lisboa"}

	fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Brasília")

	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)

	cidades[0] = "Nova Iorque"
	cidades[1] = "Londres"
	cidades[2] = "Tóquio"
	cidades[3] = "Singapura"

	fmt.Println(cidades, len(cidades), cap(cidades))

	cidades[1] = "Brasília"

	fmt.Println(cidades, len(cidades), cap(cidades))

	for indice, cidade := range cidades {

		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}

}
