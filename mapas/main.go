package main

import "fmt"
import "github.com/Zullus/cusrodego/mapas/model"

var cache map[string]model.Imovel

func main() {

    cache = make(map[string]model.Imovel, 0)

    casa := model.Imovel{}
    casa.Nome = "Casa amarela"
    casa.X = 20
    casa.Y = 25
    casa.SetValor(60000)

    cache["Casa amarela"] = casa

    fmt.Println("Há uma casa amarela no cache?")

    imovel, achou := cache["Casa amarela"]

    if achou{
        fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
    }

    apto := model.Imovel{}
    apto.Nome = "Apartamento azul"
    apto.X = 19
    apto.Y = 26
    apto.SetValor(70000)

    cache[apto.Nome] = apto

    fmt.Println("Quantos itens há no cache? ", len(cache))

    for chave, imovel := range cache{

        fmt.Printf("Chave [%s] = %+v\n\r", chave, imovel)
    }

    imovel, achou = cache["Casa amarela"]

    if achou {

        delete(cache, imovel.Nome)

    }

    fmt.Println("Quantos itens há no cache? ", len(cache))

}
