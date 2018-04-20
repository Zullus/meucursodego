package operadora

import (
    "strconv"

    "github.com/Zullus/cusrodego/pacotes/prefixo"
    )

//NomeOperadora é o nome da operadora
var NomeOperadora = "EAS Telecom"

//PrefixoDaCapitalOperadora prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora