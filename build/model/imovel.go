package model

import "errors"

//Imovel é uma struct que armagena os dados de um imóvel
type Imovel struct{
    X     int    `json:"coordanada_x"`
    Y     int    `json:"coordanada_y"`
    Nome  string    `json:"coordanada_nome"`
    valor int    //Não pode ser exportado!

}

/*
    ErrValorDeImovelInvalido erro aprensentado qdo o valor do imóvel é baixo
*/
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para o imóvel")

//ErrValorDeImovelMuitoAlto O valor do imóvel é muito alto
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

//No caso, para atrelar uma função na struct, se coloca a strucr antes do nome da função
//SetValor define valor do imóvel quando definido
func (i *Imovel) SetValor(valor int) (err error) {

    err = nil

    if valor < 50000{

        err = ErrValorDeImovelInvalido
        return

    } else if valor > 10000000 {

        err = ErrValorDeImovelMuitoAlto
        return
    }

    i.valor = valor

    return
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int{
    return i.valor
}