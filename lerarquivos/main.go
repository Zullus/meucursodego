package main

import(

    //"bufio"
    "encoding/csv"
    "fmt"
    "os"

)

func main() {

    arquivo, err := os.Open("cidades.csv")

    if err != nil{

        fmt.Println("[main]Houve um erro ao abrir o arquivo. Erro: ", err.Error())

        return
    }

    /*scanner:= bufio.NewScanner(arquivo)

    for scanner.Scan(){

        linha := scanner.Text()

        fmt.Println("O conteúdo da linhas é: ", linha)

    }*/

    leitorCsv := csv.NewReader(arquivo)

    conteudo, err := leitorCsv.ReadAll()

    if err != nil{

        fmt.Println("[main]Houve um erro ao ler o arquivo .csv. Erro: ", err.Error())

        return
    }

    for indiceLinha, linha := range conteudo{

        fmt.Printf("Linha[%d] é %s\n\r", indiceLinha, linha)

        for indiceItem, item := range linha{

            fmt.Printf("Item[%d] é %s\n\r", indiceItem, item)
        }
    }

    arquivo.Close()


}