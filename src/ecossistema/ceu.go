package ecossistema

import (
	"fmt"
)

type Ceu struct {
	ambienteC *Ambiente
}

func CeuNovo(a *Ambiente) *Ceu {
	ret := Ceu{}
	ret.ambienteC = a
	return &ret
}

func (a *Ceu) CeuValorCorrente() int {
	return int(a.ambienteC.LuzCorrenteValor() - (a.ambienteC.NuvemCorrenteValor() / 3))
}

func (a *Ceu) CeuNomeCorrente() string {
	return a.CeuLuminosidadeNome(a.CeuValorCorrente())
}

func (a *Ceu) CeuInfo() string {
	var ret string = fmt.Sprintf("Ceu :  %d - %s", a.CeuValorCorrente(), a.CeuNomeCorrente())

	return ret
}

func (a *Ceu) CeuLuminosidadeNome(_luz int) string {
	var ret string = ""

	if _luz < 10 {
		ret = "Muito Escuro !"
	}

	if _luz >= 10 && _luz < 20 {
		ret = "Escuro !"

	}

	if _luz >= 20 && _luz < 60 {
		ret = "Claro !"
	}

	if _luz >= 60 && _luz < 80 {
		ret = "Muito Claro !"
	}

	if _luz >= 80 {
		ret = "Muito Muito Claro !"
	}

	return ret
}
