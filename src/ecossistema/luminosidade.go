package ecossistema

import (
	"math/rand"
)

type luminosidade struct {
	ambienteC *Ambiente

	luz         float32
	luzcontador int
	luzlimite   int

	luzdiamax   float32
	luznoitemax float32
}

func luminosidadeNovo(a *Ambiente) *luminosidade {
	ret := luminosidade{}
	ret.ambienteC = a

	ret.luz = 0

	ret.luzdiamax = 100
	ret.luznoitemax = 18


	return &ret
}

func (a *luminosidade) claridade() {





		if a.ambienteC.fase == "Dia" {
			a.luz = float32(rand.Intn(int(a.luzdiamax))) + rand.Float32()

			if a.luz < 40 {
				a.luz += 40
			}
		}

		if a.ambienteC.fase == "Noite" {
			a.luz = float32(rand.Intn(int(a.luznoitemax))) + rand.Float32()
		}


}

func (a *luminosidade) luminosidadeNomeCorrente() string {
	return a.luminosidadeNome(a.luz)
}

func (a *luminosidade) luminosidadeNome(_luz float32) string {
	var ret string = ""

	if _luz >= 0 && _luz < 10 {
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
