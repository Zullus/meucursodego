package main

import (
    "bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
//    "log"
    "os"
    "strings"
    "time"

    "sync"

    "github.com/Zullus/cusrodego/goroutines/model"
)

var orquestrador sync.WaitGroup

func main() {

    orquestrador.Add(2) //número de tarefas

    go traduzirParaJSON("saopaulo") //excuta as tarefas em outra tread e saí do programa principal, nem os prints são exibidos

    go traduzirParaJSON("riodejaneiro") //excuta as tarefas em outra tread e saí do programa principal, nem os prints são exibidos

    orquestrador.Wait()

}

func traduzirParaJSON( nomedoArquivo string) {


    fmt.Println(time.Now(), " - Começando a truadução do arquivo: ", nomedoArquivo)

        arquivo, err := os.Open(nomedoArquivo + ".csv")

    if err != nil{

        fmt.Println("[main]Houve um erro ao abrir o arquivo. Erro: ", err.Error())

        return
    }

    defer arquivo.Close() //Lembra do compilador lembrar de fechar o arquivo

    leitorCsv := csv.NewReader(arquivo)

    conteudo, err := leitorCsv.ReadAll()

    if err != nil{

        fmt.Println("[main]Houve um erro ao ler o arquivo .csv. Erro: ", err.Error())

        return
    }

    arquivoJSON, err := os.Create(nomedoArquivo + ".json")

    if err != nil{

        fmt.Println("[main]Houve um erro ao criar o arquivo .json. Erro: ", err.Error())

        return
    }

    defer arquivoJSON.Close()

    escritor := bufio.NewWriter(arquivoJSON)

    escritor.WriteString("[\r\n")

    for _, linha := range conteudo{

        for indiceItem, item := range linha{

            dados := strings.Split(item, "/")

            cidade := model.Cidade{}
            cidade.Nome   = dados[0]
            cidade.Estado = dados[1]

            fmt.Printf("Cidade %+v\r\n", cidade)

            cidadeJSON, err := json.Marshal(cidade)

            if err != nil{

                fmt.Println("[main]Houve um erro ao gerar o JSON no item", item, ". Erro: ", err.Error())

            }

            escritor.WriteString("  " + string(cidadeJSON))

            if (indiceItem + 1) < len(linha){ //Escreve a vírgula do JSon

                escritor.WriteString(", \r\n")

            }
        }
    }

    escritor.WriteString("\r\n]")

    escritor.Flush() //Libera o buffer

    fmt.Println(time.Now(), " - A traduçaõ do arquivo ", nomedoArquivo, " foi finalizada")

    orquestrador.Done()//avisa que terminou a tarefa

}