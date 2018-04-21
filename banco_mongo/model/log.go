package model

//RegistraLog armazena quando a página foi visitada
type RegistraLog struct {
	DataVisita string `json: "dataVisita" bson: "dataVista"`
	Quem       string `json:"quem,omitempty" bson: "quem,omitempty"` //omitemppy - omite o campo se vazio
}
