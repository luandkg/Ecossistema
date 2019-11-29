package ecossistema

import (
	"fmt"
	"math/rand"
)

type Luminosidade struct {
	ambienteC *Ambiente

	luz         float32
	luzcontador int
	luzlimite   int

	luzdiamax   float32
	luznoitemax float32
}

func LuminosidadeNovo(a *Ambiente) *Luminosidade {
	ret := Luminosidade{}
	ret.ambienteC = a

	ret.luz = 20

	ret.luzdiamax = 100
	ret.luznoitemax = 18

	return &ret
}

func (a *Luminosidade) Iluminar() {

	if a.ambienteC.Fase() == "Dia" {
		a.luz = float32(rand.Intn(int(a.luzdiamax))) + rand.Float32()

		if a.luz < 40 {
			a.luz += 40
		}
	}

	if a.ambienteC.Fase() == "Noite" {
		a.luz = float32(rand.Intn(int(a.luznoitemax))) + rand.Float32()
	}

}

func (a *Luminosidade) LuminosidadeNomeCorrente() string {
	return a.LuminosidadeNome(a.luz)
}

func (a *Luminosidade) LuzCorrente() string {
	return a.LuminosidadeNome(a.luz)
}

func (a *Luminosidade) LuzCorrenteValor() float32 {
	return a.luz
}

func (a *Luminosidade) LuminosidadeNome(_luz float32) string {
	var ret string = ""

	if _luz >= 0 && _luz < 10 {
		ret = "Muito Escuro"
	}

	if _luz >= 10 && _luz < 20 {
		ret = "Escuro"
	}

	if _luz >= 20 && _luz < 60 {
		ret = "Claro"
	}

	if _luz >= 60 && _luz < 80 {
		ret = "Muito Claro"
	}

	if _luz >= 80 {
		ret = "Muitissimo Claro"
	}

	return ret
}

func (a *Luminosidade) LuzInfo() string {
	var ret string = fmt.Sprintf("Luz :  %.2f - %s", a.LuzCorrenteValor(), a.LuminosidadeNomeCorrente())

	return ret
}
