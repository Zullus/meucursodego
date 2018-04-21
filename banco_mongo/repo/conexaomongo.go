package repo

import mgo "gopkg.in/mgo.v2"

//SessaoMongo é a conexão com o banco de dados
var SessaoMongo *mgo.Session

//AbreSessaoComMongo exporta essa sessão para a aplicação
func AbreSessaoComMongo() (err error) {

	err = nil

	SessaoMongo, err = mgo.Dial("mongodb://192.168.0.10:27017/cursodego")

	if err != nil {
		return
	}

	return

}
