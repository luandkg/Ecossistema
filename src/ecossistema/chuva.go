package ecossistema



func (a *Ambiente) chover() float32 {
	var valor float32=0

	var fator1 float32 = a.vento/5
	var fator2 float32 = a.umidade/5
	var fator3 float32 = a.temperaturaCorrente
	var fator4 float32 = a.nuvem/5


	if a.temperaturaCorrente<=0 {
		fator3 = 25
	}else if a.temperaturaCorrente >=0 && a.temperaturaCorrente <20 {
		fator3=20
	}else if a.temperaturaCorrente >=20 && a.temperaturaCorrente <30 {
		fator3=10
	}else{
		fator3=0
	}

	valor = fator1+fator2+fator3+fator4

	return valor
}

func (a *Ambiente) chuvaNomeCorrente() string {
	return a.chuvaNome(a.chover())
}
func (a *Ambiente) chuvaNome(_chuva float32) string {
	var ret string = ""

	if _chuva >= 0 && _chuva < 40 {
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
