package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/Zullus/cusrodego/json_unmarshall/model"
)

func main() {

    cliente := &http.Client{ //Tem de ser um ponteiro

        Timeout: time.Second * 30,
    }

    request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)

    if err != nil{

        fmt.Println("[main] Erro ao criar um requst Servidor Json Placeholder. Erro: ", err.Error())

        return
    }

    request.SetBasicAuth("teste", "teste")

    resposta, err := cliente.Do(request)

    if err != nil{

        fmt.Println("[main] Erro ao pesquisar na página do Servidor Json Placeholder. Erro: ", err.Error())

        return
    }

    defer resposta.Body.Close()

    if resposta.StatusCode == 200{

        corpo, err := ioutil.ReadAll(resposta.Body)

        if err != nil{

            fmt.Println("[main] Erro ao ler o conteúdo da página Servidor Json Placeholder. Erro: ", err.Error())

            return
        }

        post := model.BlogPost{}

        err = json.Unmarshal(corpo, &post) //Passa o ponteiro para a struct

        if err != nil{

            fmt.Println("[main] Erro ao converter o retorno do JSon no servidor. Erro: ", err.Error())

            return
        }

        fmt.Printf("Post é: %+v\r\n", post)

    }
}