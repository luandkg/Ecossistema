package ecossistema

import "math/rand"

type nuvens struct {
	ambienteC *Ambiente

	nuvem         float32
	nuvemcontador int
	nuvemlimite   int
}

func nuvensNovo(a *Ambiente) *nuvens {
	ret := nuvens{}
	ret.ambienteC = a

	ret.nuvem = 0
	ret.nuvemcontador = 21
	ret.nuvemlimite = 20

	return &ret
}

func (a *nuvens) ceu() {

	if a.ambienteC.fase == "Dia" && a.ambienteC.fasecontador == 0 {
		a.nuvemcontador = 0
	}

	if a.ambienteC.fase == "Noite" && a.ambienteC.fasecontador == 0 {
		a.nuvemcontador = 0
	}

	if a.nuvemcontador >= a.nuvemlimite {

		a.nuvem = float32(rand.Intn(100)) + rand.Float32()

		a.nuvemcontador = 0

	} else {
		a.nuvemcontador++
	}
}

func (a *nuvens) nuvemCorrente() string {
	return a.nuvemNome(a.nuvem)
}

func (a *nuvens) nuvemNome(_nuvem float32) string {
	var ret string = ""

	if _nuvem >= 0 && _nuvem < 10 {
		ret = "Sem Nuvem"
	}

	if _nuvem >= 10 && _nuvem < 20 {
		ret = "Cirrus"
	}

	if _nuvem >= 20 && _nuvem < 30 {
		ret = "Cirrocumulus"
	}

	if _nuvem >= 30 && _nuvem < 40 {
		ret = "Cirrostratus"
	}

	if _nuvem >= 40 && _nuvem < 50 {
		ret = "Altocumulus"
	}

	if _nuvem >= 50 && _nuvem < 60 {
		ret = "Altostratus"
	}

	if _nuvem >= 60 && _nuvem < 70 {
		ret = "Stratus"
	}

	if _nuvem >= 70 {
		ret = "Stratocumulus"
	}

	return ret
}
