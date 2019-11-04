package ecossistema

import (
	"fmt"
	"math/rand"
)

type Nuvens struct {
	ambienteC *Ambiente

	nuvem            float32
	nuvemcontador    int
	nuvemlimite      int
	nuvemfinalizador int
}

func NuvensNovo(a *Ambiente) *Nuvens {
	ret := Nuvens{}
	ret.ambienteC = a

	ret.nuvem = 0
	ret.nuvemcontador = 21
	ret.nuvemlimite = 20
	ret.nuvemfinalizador = ret.nuvemlimite

	return &ret
}

func (a *Nuvens) Nublar() {

	a.nuvemcontador++

	if a.nuvemcontador >= a.nuvemfinalizador {

		a.nuvem = float32(rand.Intn(99)) + rand.Float32()

		a.nuvemcontador = 0
		a.nuvemfinalizador = a.nuvemlimite + rand.Intn(30)

		//utils.Log("logs/ambientelimitador.txt", "Nuvem - "+strconv.Itoa(a.nuvemfinalizador) + "   " + a.NuvemNomeCorrente())

	}
}

func (a *Nuvens) NuvemNomeCorrente() string {
	return a.NuvemNome(a.nuvem)
}

func (a *Nuvens) NuvemCorrente() string {
	return a.NuvemNome(a.nuvem)
}
func (a *Nuvens) NuvemCorrenteValor() float32 {
	return a.nuvem
}

func (a *Nuvens) NuvemNome(_nuvem float32) string {
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

func (a *Nuvens) NuvemInfo() string {
	var ret string = fmt.Sprintf("Nuvens :  %.2f - %s", a.NuvemCorrenteValor(), a.NuvemNomeCorrente())

	return ret
}
