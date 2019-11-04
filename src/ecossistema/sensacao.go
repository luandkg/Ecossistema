package ecossistema

type Sensacao struct {
	ambienteC *Ambiente
}

func SensacaoNovo(a *Ambiente) *Sensacao {
	ret := Sensacao{}
	ret.ambienteC = a
	return &ret
}


func (a *Sensacao) SensacaoInfo() string {
	return "Sensação Térmica : " + a.Sensacao()
}

func (a *Sensacao) Sensacao() string {
	var ret string = ""

	if a.ambienteC.TemperaturaCorrente() >= 20 && a.ambienteC.TemperaturaCorrente() <= 28 && a.ambienteC.UmidadeCorrenteValor() >= 40 && a.ambienteC.UmidadeCorrenteValor() <= 60 {
		ret = "Agradável"
	}

	if a.ambienteC.TemperaturaCorrente() >= 20 && a.ambienteC.TemperaturaCorrente() <= 28 && a.ambienteC.UmidadeCorrenteValor() < 40 {
		ret = "Umidade Baixa"
	}

	if a.ambienteC.TemperaturaCorrente() >= 20 && a.ambienteC.TemperaturaCorrente() <= 28 && a.ambienteC.UmidadeCorrenteValor() > 60 {
		ret = "Umidade Alta"
	}

	if a.ambienteC.TemperaturaCorrente() < 20 && a.ambienteC.UmidadeCorrenteValor() >= 40 && a.ambienteC.UmidadeCorrenteValor() <= 60 {
		ret = "Frio"
	}

	if a.ambienteC.TemperaturaCorrente() > 28 && a.ambienteC.UmidadeCorrenteValor() >= 40 && a.ambienteC.UmidadeCorrenteValor() <= 60 {
		ret = "Quente"
	}

	if a.ambienteC.TemperaturaCorrente() < 20 && a.ambienteC.UmidadeCorrenteValor() < 40 {
		ret = "Frio e Seco"
	}

	if a.ambienteC.TemperaturaCorrente() < 20 && a.ambienteC.UmidadeCorrenteValor() > 60 {
		ret = "Frio e Umido"
	}

	if a.ambienteC.TemperaturaCorrente() > 28 && a.ambienteC.UmidadeCorrenteValor() < 40 {
		ret = "Quente e Seco"
	}

	if a.ambienteC.TemperaturaCorrente() > 28 && a.ambienteC.UmidadeCorrenteValor() > 60 {
		ret = "Quente e Umido"
	}

	return ret
}

