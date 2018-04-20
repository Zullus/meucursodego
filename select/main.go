package main

import(

    "fmt"
    "time"

)

var irc = make (chan string)
var sms = make (chan string)

func eai(canal chan string) {

    for{
        canal <- "eai" //Ponger escreve no channel canal
    }


}

func blz(canal chan string) {

    for{
        canal <- "blz" //Ponger escreve no channel canal
    }


}

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

func impressora() {

    var msg string

    for{

        select {

            case msg = <-irc:

                fmt.Println(msg)

                time.Sleep(time.Second * 1)

            case msg = <-sms:

                fmt.Println("Zzz....zzZ ", msg)

                time.Sleep(time.Second * 1)

            }
    }

}

func main() {

    go pinger(irc)

    go ponger(irc)

    go eai(sms)

    go blz(sms)

    go impressora()

    var entrada string

    fmt.Scanln(&entrada)


}