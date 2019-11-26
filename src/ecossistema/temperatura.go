package ecossistema

import (
	"fmt"
	"math/rand"
)

type Temperatura struct {
	ambienteC *Ambiente

	tempbase float32

	Corrente float32

	temp1 float32
	temp2 float32
	temp3 float32
	temp4 float32

	tempdiamin float32
	tempdiamax float32

	tempmin float32
	tempmax float32
}

func TemperaturaNovo(a *Ambiente) *Temperatura {
	ret := Temperatura{}
	ret.ambienteC = a

	ret.tempbase = 25
	ret.Corrente = 0

	ret.tempdiamin = 0
	ret.tempdiamax = 0

	ret.tempmin = 0
	ret.tempmax = 0

	return &ret
}

func (a *Temperatura) Esquentar() {
	a.TemperaturaDia()
	a.TemperaturaNoite()
}

func (a *Temperatura) TemperaturaCorrente() float32 {
	return a.Corrente
}

func (a *Temperatura) TemperaturaMedia() float32 {
	return (a.temp1 + a.temp2 + a.temp3 + a.temp4) / 4
}

func (a *Temperatura) TemperaturaDia() {

	if a.ambienteC.Fase() == "Dia" && a.ambienteC.FaseContador() == 0 {
		a.temp1 = 0
		a.temp2 = 0
		a.temp3 = 0
		a.temp4 = 0
		a.Corrente = 0

		a.tempdiamin = 0
		a.tempdiamax = 0

		if rand.Intn(100) < 50 {
			a.temp1 = a.tempbase + float32(rand.Intn(5)) + rand.Float32()
		} else {
			a.temp1 = a.tempbase - float32(rand.Intn(5)) + rand.Float32()
		}

		a.tempdiamax = a.temp1
		a.tempdiamin = a.temp1
		a.Corrente = a.temp1

	}

	if a.ambienteC.Fase() == "Dia" && a.ambienteC.FaseContador() == 30 {

		if rand.Intn(100) < 50 {
			a.temp2 = a.temp1 + float32(rand.Intn(5)) + rand.Float32()
		} else {
			a.temp2 = a.temp1 - float32(rand.Intn(5)) + rand.Float32()
		}

		if a.temp2 > a.tempdiamax {
			a.tempdiamax = a.temp2
		}
		if a.temp2 < a.tempdiamin {
			a.tempdiamin = a.temp2
		}
		a.Corrente = a.temp2
	}

	if a.ambienteC.Fase() == "Dia" && a.ambienteC.FaseContador() == 60 {

		if rand.Intn(100) < 50 {
			a.temp3 = a.temp2 + float32(rand.Intn(5)) + rand.Float32()
		} else {
			a.temp3 = a.temp2 - float32(rand.Intn(5)) + rand.Float32()
		}

		if a.temp3 > a.tempdiamax {
			a.tempdiamax = a.temp3
		}
		if a.temp3 < a.tempdiamin {
			a.tempdiamin = a.temp3
		}

		a.Corrente = a.temp3
	}

	if a.ambienteC.Fase() == "Dia" && a.ambienteC.FaseContador() == 90 {

		if rand.Intn(100) < 50 {
			a.temp4 = a.temp3 + float32(rand.Intn(5)) + rand.Float32()
		} else {
			a.temp4 = a.temp3 - float32(rand.Intn(5)) + rand.Float32()
		}

		if a.temp4 > a.tempdiamax {
			a.tempdiamax = a.temp4
		}
		if a.temp4 < a.tempdiamin {
			a.tempdiamin = a.temp4
		}
		a.Corrente = a.temp4
	}

}

func (a *Temperatura) TemperaturaNoite() {

	if a.ambienteC.Fase() == "Noite" && a.ambienteC.FaseContador() == 0 {
		a.temp1 = 0
		a.temp2 = 0
		a.temp3 = 0
		a.temp4 = 0
		a.Corrente = 0

		var reduz int = rand.Intn(5)

		if rand.Intn(100) < 50 {
			a.temp1 = (a.tempbase + float32(rand.Intn(5))) - float32(reduz) + rand.Float32()
		} else {
			a.temp1 = (a.tempbase - float32(rand.Intn(5))) - float32(reduz) + rand.Float32()
		}

		a.tempdiamax = a.temp1
		a.tempdiamin = a.temp1
		a.Corrente = a.temp1
	}

	if a.ambienteC.Fase() == "Noite" && a.ambienteC.FaseContador() == 30 {

		if rand.Intn(100) < 50 {
			a.temp2 = a.temp1 + float32(rand.Intn(5)) + rand.Float32()
		} else {
			a.temp2 = a.temp1 - float32(rand.Intn(5)) + rand.Float32()
		}

		if a.temp2 > a.tempdiamax {
			a.tempdiamax = a.temp2
		}
		if a.temp2 < a.tempdiamin {
			a.tempdiamin = a.temp2
		}
		a.Corrente = a.temp2
	}

	if a.ambienteC.Fase() == "Noite" && a.ambienteC.FaseContador() == 60 {

		if rand.Intn(100) < 50 {
			a.temp3 = a.temp2 + float32(rand.Intn(5)) + rand.Float32()
		} else {
			a.temp3 = a.temp2 - float32(rand.Intn(5)) + rand.Float32()
		}

		if a.temp3 > a.tempdiamax {
			a.tempdiamax = a.temp3
		}
		if a.temp3 < a.tempdiamin {
			a.tempdiamin = a.temp3
		}
		a.Corrente = a.temp3
	}

	if a.ambienteC.Fase() == "Noite" && a.ambienteC.FaseContador() == 90 {

		if rand.Intn(100) < 50 {
			a.temp4 = a.temp3 + float32(rand.Intn(5)) + rand.Float32()
		} else {
			a.temp4 = a.temp3 - float32(rand.Intn(5)) + rand.Float32()
		}

		if a.temp4 > a.tempdiamax {
			a.tempdiamax = a.temp4
		}
		if a.temp4 < a.tempdiamin {
			a.tempdiamin = a.temp4
		}
		a.Corrente = a.temp4
	}

}

func (a *Temperatura) TemperaturaModo(_temp float32) string {
	var ret string = ""

	if _temp < 10 {
		ret = "Muito Frio !"
	}

	if _temp >= 10 && _temp < 20 {
		ret = "Frio !"
	}

	if _temp >= 20 && _temp < 28 {
		ret = "Normal !"
	}

	if _temp >= 28 && _temp < 32 {
		ret = "Quente !"
	}

	if _temp >= 32 {
		ret = "Muito Quente !"
	}

	return ret
}

func (a *Temperatura) TemperaturaInfo() string {
	return "Temperatura - " + fmt.Sprintf("%.2f ºC", a.TemperaturaCorrente()) + " : " + a.TemperaturaModo(a.TemperaturaCorrente())
}

func (a *Temperatura) Kelvin(valor float32) float32 {
	return valor + 273.15
}

func (a *Temperatura) Fahrenheit(valor float32) float32 {
	return ((valor * 9) / 5) + 32
}

func (a *Temperatura) Rankine(valor float32) float32 {
	return (((valor * 9) / 5) + 32) + 459.67
}
