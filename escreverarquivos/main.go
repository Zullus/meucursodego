package main

import(
    "bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "strings"

    "github.com/Zullus/cusrodego/escreverarquivos/model"
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

    arquivoJSON, err := os.Create("cidades.json")

    if err != nil{

        fmt.Println("[main]Houve um erro ao criar o arquivo .json. Erro: ", err.Error())

        return
    }

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

    arquivoJSON.Close()

    arquivo.Close()


}