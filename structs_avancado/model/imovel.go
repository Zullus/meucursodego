package model

//Imovel é uma struct que armagena os dados de um imóvel
type Imovel struct{
    X     int    `json:"coordanada_x"`
    Y     int    `json:"coordanada_y"`
    Nome  string    `json:"coordanada_nome"`
    valor int    //Não pode ser exportado!

}

//No caso, para atrelar uma função na struct, se coloca a strucr antes do nome da função
//SetValor define valor do imóvel quando definido
func (i *Imovel) SetValor(valor int) {
    i.valor = valor
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int{
    return i.valor
}