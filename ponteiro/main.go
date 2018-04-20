package main

import "fmt"

//Imovel é uma struct que armagena os dados de um imóvel
type Imovel struct{
    X     int
    Y     int
    Nome  string
    valor int
}

func main() {

    casa := new(Imovel) //Aponta para um ponteiro
    fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

    chacara := Imovel{17, 28, "Chácara Linda", 280000}
    fmt.Printf("Chácara é: %p - %+v\r\n", &chacara, chacara)

    mudaImovel(&chacara) //Passa o endereço de memória da estrura
    fmt.Printf("Chácara é: %p - %+v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
/*No Go ele faz uma cópia dos dados quando passa para uma variável, com ponteiro, ele não faz a cópia
*/

    imovel.X = imovel.X + 10
    imovel.Y = imovel.Y - 5

}