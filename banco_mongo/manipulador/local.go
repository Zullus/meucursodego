package manipulador

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Zullus/cursodego/banco_mongo/repo"

	"github.com/Zullus/cursodego/banco_mongo/model"
)

//Local é a ligação com o template HTML com o banco de dados
func Local(w http.ResponseWriter, r *http.Request) {

	local := model.Local{}

	//Paramentro para a query, pegando da URL
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:]) //depois do sétimo caracter -> Atoi converte para string

	if err != nil {

		http.Error(w, "Não foi enviado um número válido!", http.StatusBadRequest)

		fmt.Println("[Local] Erro ao converter o número enviado: ", err.Error())
	}

	//Query

	sql := "SELECT country, city, telcode FROM place WHERE telcode = ?" //Char coringa
	/*
	   É fundamental que o SELECT tenha os mesmos elementos da struct da model, nem a mais e nem a menos!!!!
	*/
	linha, err := repo.Db.Queryx(sql, codigoTelefone) //Executa a query

	if err != nil {

		http.Error(w, "Não possível pesquisar esse número!", http.StatusInternalServerError)

		fmt.Println("[Local] Não foi possível executar a query: ", sql, " Erro: ", err.Error())

		return
	}

	//Interar as linhas do banco

	for linha.Next() {

		err = linha.StructScan(&local)

		if err != nil {

			http.Error(w, "Não possível exibir esse código!", http.StatusInternalServerError)

			fmt.Println("[Local] Não foi fazer o biding dos dados do banco na struct: ", err.Error())

			return
		}
	}
	//O erro pode ser fora do for

	//MondoDB
	localMongo, err := repo.ObtemLocal(codigoTelefone)

	if err != nil {

		http.Error(w, "Não possível pesquisar esse número!", http.StatusInternalServerError)

		fmt.Println("[Local] Não foi possível executar no MongoDB. Erro: ", err.Error())

		return
	}

	//if err := ModelosLocal.ExecuteTemplate(w, "local.html", local); err != nil { //Busca no MySQL
	if err := ModelosLocal.ExecuteTemplate(w, "local.html", localMongo); err != nil { //Busca no MongoDB

		http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError) //Não exibe a função ou detalher pq será exposto

		fmt.Println("[Local] Erro na execução do modelo: ", err.Error()) //Exibe a função por ser um erro interno

	}

	//Inserir dados no banco, pode ser update ou detele também
	sql = "INSERT INTO logquery (daterequest) VALUES (?)"

	resultado, err := repo.Db.Exec(sql, time.Now().Format("02/01/2006 15:04:05"))

	if err != nil {

		fmt.Println("[Local] Erro na inclusão do log: ", sql, " - Erro: ", err.Error())

	}

	linhasAfetadas, err := resultado.RowsAffected()

	if err != nil {

		fmt.Println("[Local] Erro ao pegar o número de linhas afetadas pelo comaando: ", sql, " - Erro: ", err.Error())

	}

	log := model.RegistraLog{}
	log.DataVisita = time.Now().Format("02/01/2006 15:04:05")

	err = repo.SalvaLog(log)

	if err != nil {

		fmt.Println("[local] Erro na inclusão do log no MongoDB", err.Error())

	}

	println("Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")

}
