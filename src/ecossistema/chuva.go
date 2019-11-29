package ecossistema

import (
	"fmt"
	"math/rand"
)

type Chuva struct {
	ambienteC *Ambiente

	chovenotipo         string
	chovendoespecial    string
	chovendointensidade float32
}

func ChuvaNovo(a *Ambiente) *Chuva {
	ret := Chuva{}
	ret.ambienteC = a
	ret.chovenotipo = " - "
	ret.chovendoespecial = " - "
	ret.chovendointensidade = 0

	return &ret
}

func (a *Chuva) ChuvaTipo() string {
	return a.chovenotipo
}

func (a *Chuva) ChuvaModo() string {
	return a.chovendoespecial
}

func (a *Chuva) Chover() {
	a.ChuvaCorrenteValor()
}

func (a *Chuva) ChuvaCorrenteValor() float32 {
	var valor float32 = 0

	var fator1 float32 = a.ambienteC.vento / 4
	var fator2 float32 = a.ambienteC.umidade / 4
	var fator3 float32 = a.ambienteC.TemperaturaCorrente()
	var fator4 float32 = a.ambienteC.nuvem / 4

	if a.ambienteC.TemperaturaCorrente() <= 0 {
		fator3 = 25
	} else if a.ambienteC.TemperaturaCorrente() >= 0 && a.ambienteC.TemperaturaCorrente() < 20 {
		fator3 = 20
	} else if a.ambienteC.TemperaturaCorrente() >= 20 && a.ambienteC.TemperaturaCorrente() < 30 {
		fator3 = 10
	} else {
		fator3 = 0
	}

	valor = fator1 + fator2 + fator3 + fator4

	if valor >= 50 {

		if a.chovenotipo == " - " {

			var chovendoint int = rand.Intn(100)
			if chovendoint < 25 {
				a.chovendointensidade = float32(rand.Intn(int(2))) + rand.Float32()
			}

			if chovendoint >= 25 && chovendoint < 50 {
				a.chovendointensidade = 2.5 + float32(rand.Intn(int(9))) + rand.Float32()
			}

			if chovendoint >= 50 && chovendoint < 75 {
				a.chovendointensidade = 10 + float32(rand.Intn(int(40))) + rand.Float32()
			}

			if chovendoint >= 75 {
				a.chovendointensidade = 50 + float32(rand.Intn(int(50))) + rand.Float32()
			}

			var chuvatiponum int = rand.Intn(100)
			if chuvatiponum >= 0 && chuvatiponum < 50 {
				a.chovenotipo = "Frontal"
			} else if chuvatiponum >= 50 && chuvatiponum < 75 {
				a.chovenotipo = "Orográfica"
			} else if chuvatiponum >= 75 {
				a.chovenotipo = "Convectiva"
			}

		}

		if a.chovendoespecial == " - " {
			var especial int = rand.Intn(100)
			if especial >= 0 && especial < 50 {
				a.chovendoespecial = "Água"
			} else if especial >= 50 && especial < 75 {
				a.chovendoespecial = "Granizo"
			} else if especial >= 75 {
				a.chovendoespecial = "Neve"
			}

		}

	} else {
		a.chovenotipo = " - "
		a.chovendoespecial = " - "
		a.chovendointensidade = 0
	}

	return valor
}

func (a *Chuva) ChuvaNomeCorrente() string {
	return a.ChuvaNome(a.ChuvaCorrenteValor())
}

func (a *Chuva) ChuvaCorrente() string {
	return a.ChuvaNome(a.ChuvaCorrenteValor())
}

func (a *Chuva) ChuvaNome(_chuva float32) string {
	var ret string = ""

	if _chuva < 30 {
		ret = "Sem Chuva"
	}

	if _chuva >= 30 && _chuva < 40 {
		ret = "Neblina"
	}

	if _chuva >= 40 && _chuva < 50 {
		ret = "Nevoeiro"
	}

	if _chuva >= 50 && _chuva < 60 {
		ret = "Chuvisco"
	}

	if _chuva >= 60 && _chuva < 70 {
		ret = "Chuva"
	}

	if _chuva >= 80 {
		ret = "Chuva Forte"
	}

	return ret
}

func (a *Chuva) ChuvaInfo() string {
	var ret string = fmt.Sprintf("Chuva :  %.2f - %s", a.ChuvaCorrenteValor(), a.ChuvaNomeCorrente())

	return ret
}

func (a *Chuva) ChuvaIntensidade() float32 {
	return a.chovendointensidade
}

func (a *Chuva) ChuvaIntensidadeInfo() string {
	if a.ChuvaCorrenteValor() >= 50 {
		return fmt.Sprintf("%.2f mmh", a.ChuvaIntensidade())
	} else {
		return " - "
	}
}

func (a *Chuva) ChuvaIntensidadeStatus() string {

	var ret string = ""

	if a.ChuvaCorrenteValor() >= 50 {

		if a.chovendointensidade < 2.5 {
			ret = "Fraca"
		}

		if a.chovendointensidade >= 2.5 && a.chovendointensidade < 10 {
			ret = "Moderada"
		}

		if a.chovendointensidade >= 10 && a.chovendointensidade < 50 {
			ret = "Forte"
		}

		if a.chovendointensidade > 50 {
			ret = "Violenta"
		}

	} else {
		ret = " - "
	}

	return ret

}
