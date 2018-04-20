package manipulador

import "html/template"

//Modelos é armazena os modelos de página executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal aramazena a página local com o banco de ddos
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))