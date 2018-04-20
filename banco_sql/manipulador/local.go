package manipulador

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/Zullus/cusrodego/servidor_web/model"
)

//Local é a conexao com o template
func Local(w http.ResponseWriter, r *http.Request) {

    local := model.Local{}

    codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])

    if err != nil{

        http.Error(w, "Não foi enviado um número válio", http.StatusBadRequest) //Não exibe a função ou detalher pq será exposto

        fmt.Println("[Local] Erro ao converter o número enviado: ", err.Error()) //Exibe a função por ser um erro interno

    }

    sql := "SELECT contry, city, telcode FROM place WHERE telcode = ?"

    linha, err := repo.Db.Qureryx(sql, codigoTelefone)

    if err != nil{

        http.Error(w, "Não foi possível pesquisar esse número banco de dados", http.StatusInternalServerError) //Não exibe a função ou detalher pq será exposto

        fmt.Println("[Local] Não possível executar a query: ", sql, " - Erro: ", err.Error()) //Exibe a função por ser um erro interno

        return

    }

    for linha.Next(){

        err = linha.StructScan(&local)

        if err != nil{

            http.Error(w, "Não foi possível pesquisar esse número banco de dados", http.StatusInternalServerError) //Não exibe a função ou detalher pq será exposto

            fmt.Println("[Local] Não possível fazer o binding dos dados na struct: Erro: ", err.Error()) //Exibe a função por ser um erro interno

        }

    }


    if err := Modelos.ExecuteTemplate(w, "local.html", pagina); err != nil{

        http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError) //Não exibe a função ou detalher pq será exposto

        fmt.Println("[Local] Erro na execução do modelo: ", err.Error()) //Exibe a função por ser um erro interno

    }

}