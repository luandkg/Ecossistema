package ecossistema

import "fmt"

type Chuva struct {
	ambienteC *Ambiente
}

func ChuvaNovo(a *Ambiente) *Chuva {
	ret := Chuva{}
	ret.ambienteC = a
	return &ret
}

func (a *Chuva) Chover() float32 {
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

	return valor
}

func (a *Chuva) ChuvaNomeCorrente() string {
	return a.ChuvaNome(a.Chover())
}

func (a *Chuva) ChuvaCorrenteValor() float32 {
	return (a.Chover())
}

func (a *Chuva) ChuvaCorrente() string {
	return a.ChuvaNome(a.Chover())
}

func (a *Chuva) ChuvaNome(_chuva float32) string {
	var ret string = ""

	if _chuva < 40 {
		ret = "Sem Chuva"
	}

	if _chuva >= 40 && _chuva < 50 {
		ret = "Neblina"
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
	var ret string = fmt.Sprintf("Chuva :  %.2f - %s", a.Chover(), a.ChuvaNomeCorrente())

	return ret
}
