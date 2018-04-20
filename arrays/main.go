package main

import (
	"fmt"
)

func main() {

	var teste [3]int

	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Println("Qual o tamanho do array?", len(teste))

	cantores := [2]string{"Michael Jackson", "Jonh Lennon"}
	fmt.Printf("O que há nesse array?\n\r%+v\n\r", cantores)

	capitais := [...]string{"Lisboa", "Brasília", "Luanda", "Maputo"}
	fmt.Println("Qual o tamanho do array?", len(capitais))

	for indice, cidade := range capitais {

		fmt.Printf("Capitais[%d] é %s \n\r", indice, cidade)
	}

}
