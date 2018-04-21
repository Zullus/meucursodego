package repo

import (
	"github.com/Zullus/cursodego/banco_mongo/model"
	"gopkg.in/mgo.v2/bson"
)

//ObtemLocal retorna um local do MongoDB
func ObtemLocal(codigoTelefone int) (local model.Local, err error) {

	sessao := SessaoMongo.Copy() //Copia a sessão do Mongo. Não copiar, pois dá pau

	defer SessaoMongo.Close()

	colecao := sessao.DB("cursodego").C("local")

	err = colecao.Find(bson.M{"telcode": codigoTelefone}).One(&local) //Faz a query. Passa o ponreiro para o mongo

	return
}

//SalvaLog Registra a consulta ao local no MongoDB
func SalvaLog(reg model.RegistraLog) (err error) {

	sessao := SessaoMongo.Copy() //Copia a sessão do Mongo. Não copiar, pois dá pau

	defer SessaoMongo.Close()

	colecao := sessao.DB("cursodego").C("logvisitas")

	err = colecao.Insert(reg)

	return

}
