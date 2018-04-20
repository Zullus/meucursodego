package main

import "fmt"

func main() {

	//var nuns []int

	//fmt.Println(nuns, len(nuns), cap(nuns))

	//nuns = make([]int, 5) //inicia um slice

	//fmt.Println(nuns, len(nuns), cap(nuns))

	capitais := []string{"Lisboa"}

	//fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Brasília")

	//fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 5)

	cidades[0] = "Nova Iorque"
	cidades[1] = "Londres"
	cidades[2] = "Madeira"
	cidades[3] = "Tóquio"
	cidades[4] = "Singapura"

	//fmt.Println(cidades, len(cidades), cap(cidades))

	cidades[1] = "Brasília"

	//fmt.Println(cidades, len(cidades), cap(cidades))

	//for indice, cidade := range cidades {

	//fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	//}

	capitaisAsia := cidades[3:5] //primeiro intem começa a contar de zero e o segundo de 1

	fmt.Println(capitaisAsia, len(capitaisAsia), cap(capitaisAsia))

	temp1 := cidades[:2] //pega os dois primeiros itens do slice ~ o ZERO é ignorado

	fmt.Println(temp1, len(temp1), cap(temp1))

	temp2 := cidades[len(cidades)-2:] //Até o fim

	fmt.Println(temp2)

	//Remove item do slice, tem de fazer com um slice temporário

	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	copy(cidades, temp)
	fmt.Println(cidades)

}
