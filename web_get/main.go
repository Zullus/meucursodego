package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func main() {

    cliente := &http.Client{ //Tem de ser um ponteiro

        Timeout: time.Second * 30,
    }

    resposta, err := cliente.Get("https://google.com.br")//get já tem request

    if err != nil{

        fmt.Println("[main] Erro ao pesquisar na página do Google Brasil. Erro: ", err.Error())

        return
    }

    defer resposta.Body.Close()

    if resposta.StatusCode == 200{

        corpo, err := ioutil.ReadAll(resposta.Body)

        if err != nil{

            fmt.Println("[main] Erro ao ler o conteúdo da página Google Brasil. Erro: ", err.Error())

            return
        }

        fmt.Println(string(corpo))
    }

    //Acesso com senha

    fmt.Println("__________________________________________________________________\n\r")

    request, err := http.NewRequest("GET", "https://google.com.br", nil)

    if err != nil{

        fmt.Println("[main] Erro ao criar um requst Google Brasil. Erro: ", err.Error())

        return
    }

    request.SetBasicAuth("teste", "teste")

    resposta, err = cliente.Do(request)

    if err != nil{

        fmt.Println("[main] Erro ao pesquisar na página do Google Brasil. Erro: ", err.Error())

        return
    }

    defer resposta.Body.Close()

    if resposta.StatusCode == 200{

        corpo, err := ioutil.ReadAll(resposta.Body)

        if err != nil{

            fmt.Println("[main] Erro ao ler o conteúdo da página Google Brasil. Erro: ", err.Error())

            return
        }

        fmt.Println(string(corpo))
    }
}