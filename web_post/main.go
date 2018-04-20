package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/Zullus/cusrodego/web_post/model"
)

func main() {

    cliente := &http.Client{ //Tem de ser um ponteiro

        Timeout: time.Second * 30,
    }

    usuario := model.Usuario{}
    usuario.ID = 1
    usuario.Nome = "Edson Santos"

    conteudoEnviar, err := json.Marshal(usuario)

    if err != nil{

        fmt.Println("[main] Erro ao Criar usuário json. Erro: ", err.Error())

        return
    }

    request, err := http.NewRequest("POST", "http://requestbin.fullcontact.com/1bx2h9r1", bytes.NewBuffer(conteudoEnviar))

    if err != nil{

        fmt.Println("[main] Erro ao criar um requst Request Bin. Erro: ", err.Error())

        return
    }

    request.SetBasicAuth("fizz", "buzz")

    request.Header.Set("content-type", "application/json; charset=utf-8")

    resposta, err := cliente.Do(request)

    if err != nil{

        fmt.Println("[main] Erro ao chamar o serviço do Request Bin. Erro: ", err.Error())

        return
    }

    defer resposta.Body.Close()

    if resposta.StatusCode == 200{

        corpo, err := ioutil.ReadAll(resposta.Body)

        if err != nil{

            fmt.Println("[main] Erro ao ler o retornado ao Request Bin. Erro: ", err.Error())

            return
        }

        fmt.Println(string(corpo))
    }
}