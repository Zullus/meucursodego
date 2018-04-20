package manipulador

import (
    "fmt"
    "net/http"
)

//Funcao é o manipulador de requisicao  HTTP
func Funcao(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintln(w, "Aqui é o manipulador udando função num pacote")

}