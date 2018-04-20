package main

import (
    "fmt"
    "encoding/json"
    "github.com/Zullus/cusrodego/erro/model"
    )

func main() {

    casa := model.Imovel{}
    casa.Nome = "Casa da Lucy"
    casa.X = 17
    casa.Y = 29
    if err:= casa.SetValor(100); err != nil{

        fmt.Println("[main] Houve um erro de atribuição do valor da casa: ", err.Error())

        if err == model.ErrValorDeImovelMuitoAlto{

            fmt.Println("Corretor, por favor verifique a sua avaliação")
        }

        return
    }

    fmt.Printf("A casa é %+v\n\r", casa)

    objJSON, err := json.Marshal(casa)

    if err != nil {

        fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error())
        return
    }

    fmt.Println("A casa em JSON: ", string(objJSON))

}
