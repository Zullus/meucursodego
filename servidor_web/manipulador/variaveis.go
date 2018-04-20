package manipulador

import "html/template"

//Modelos é armazena os modelos de página executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))