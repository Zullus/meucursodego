package main

import(

    "fmt"
    "time"

)

func pinger(canal chan string) {

    for {
        canal <- "ping" //Pinger escrecve no channel canal
    }

}

func ponger(canal chan string) {

    for{
        canal <- "pong" //Ponger escreve no channel canal
    }

}

func impressora(canal chan string) {

    for{

        msg := <-canal //msg recebe informação do canal

        fmt.Println(msg)

        time.Sleep(time.Second * 1)

    }

}

func main() {

    var canal chan string //chan -> cria channel e pode passar qqer tipo primitivo entre as GoRoutines

    canal = make(chan string)

    go pinger(canal)

    go ponger(canal)

    go impressora(canal)

    var entrada string

    fmt.Scanln(&entrada)


}