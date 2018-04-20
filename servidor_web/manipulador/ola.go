package manipulador

import (
    "fmt"
    "net/http"
    "time"

    "github.com/Zullus/cusrodego/servidor_web/model"
)

//Ola é o manipulador de requisicao  HTTP acom saida no template
func Ola(w http.ResponseWriter, r *http.Request) {

    hora := time.Now().Format("15:04:05")

    pagina := model.Pagina{}
    pagina.Hora = hora

    if err := Modelos.ExecuteTemplate(w, "ola.html", pagina); err != nil{

        http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError) //Não exibe a função ou detalher pq será exposto

        fmt.Println("[Ola] Erro na execução do modelo: ", err.Error()) //Exibe a função por ser um erro interno

    }

}