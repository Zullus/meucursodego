package model

//Local armazerna os dados da localidade pelo seu código telefônico
type Local struct {
	Pais             string `json:"pais" db:"country" bson:"country"`
	Cidade           string `json:"cidade" db:"city" bson:"city"` //Por causa do Mongo, tem de alterar como string e tratar no código para não voltar nulo
	CodigoTelefonico int    `json:"cod_telefone" db:"telcode"  bson:"telcode"`
}
