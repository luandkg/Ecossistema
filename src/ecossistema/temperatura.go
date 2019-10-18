package ecossistema

import "math/rand"

type temperatura struct {
	ambienteC *Ambiente

	tempbase float32

	temperaturaCorrente float32

	temp1 float32
	temp2 float32
	temp3 float32
	temp4 float32

	tempdiamin float32
	tempdiamax float32

	tempmin float32
	tempmax float32
}

func temperaturaNovo(a *Ambiente) *temperatura {
	ret := temperatura{}
	ret.ambienteC = a

	ret.tempbase = 25
	ret.temperaturaCorrente = 0

	ret.tempdiamin = 0
	ret.tempdiamax = 0

	ret.tempmin = 0
	ret.tempmax = 0

	return &ret
}

func (a *temperatura) temperaturaMedia() float32 {
	return (a.temp1 + a.temp2 + a.temp3 + a.temp4) / 4
}

func (a *temperatura) temperaturaDia() {

	if a.ambienteC.fase == "Dia" && a.ambienteC.fasecontador == 0 {
		a.temp1 = 0
		a.temp2 = 0
		a.temp3 = 0
		a.temp4 = 0
		a.temperaturaCorrente = 0

		a.tempdiamin = 0
		a.tempdiamax = 0

		if rand.Intn(100) < 50 {
			a.temp1 = a.tempbase + float32(rand.Intn(5)) + rand.Float32()
		} else {
			a.temp1 = a.tempbase - float32(rand.Intn(5)) + rand.Float32()
		}

		a.tempdiamax = a.temp1
		a.tempdiamin = a.temp1
		a.temperaturaCorrente = a.temp1

	}

	if a.ambienteC.fase == "Dia" && a.ambienteC.fasecontador == 30 {

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
		a.temperaturaCorrente = a.temp2
	}

	if a.ambienteC.fase == "Dia" && a.ambienteC.fasecontador == 60 {

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

		a.temperaturaCorrente = a.temp3
	}

	if a.ambienteC.fase == "Dia" && a.ambienteC.fasecontador == 90 {

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
		a.temperaturaCorrente = a.temp4
	}

}

func (a *temperatura) temperaturaNoite() {

	if a.ambienteC.fase == "Noite" && a.ambienteC.fasecontador == 0 {
		a.temp1 = 0
		a.temp2 = 0
		a.temp3 = 0
		a.temp4 = 0
		a.temperaturaCorrente = 0

		var reduz int = rand.Intn(5)

		if rand.Intn(100) < 50 {
			a.temp1 = (a.tempbase + float32(rand.Intn(5))) - float32(reduz) + rand.Float32()
		} else {
			a.temp1 = (a.tempbase - float32(rand.Intn(5))) - float32(reduz) + rand.Float32()
		}

		a.tempdiamax = a.temp1
		a.tempdiamin = a.temp1
		a.temperaturaCorrente = a.temp1
	}

	if a.ambienteC.fase == "Noite" && a.ambienteC.fasecontador == 30 {

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
		a.temperaturaCorrente = a.temp2
	}

	if a.ambienteC.fase == "Noite" && a.ambienteC.fasecontador == 60 {

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
		a.temperaturaCorrente = a.temp3
	}

	if a.ambienteC.fase == "Noite" && a.ambienteC.fasecontador == 90 {

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
		a.temperaturaCorrente = a.temp4
	}

}
