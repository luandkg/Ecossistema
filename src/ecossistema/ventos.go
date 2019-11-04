package ecossistema

import (
	"fmt"
	"math/rand"
)

type Ventos struct {
	ambienteC *Ambiente

	vento            float32
	ventocontador    int
	ventolimite      int
	ventofinalizador int

	ventoorigem  string
	ventodestino string
}

func VentosNovo(a *Ambiente) *Ventos {
	ret := Ventos{}
	ret.ambienteC = a

	ret.vento = 0
	ret.ventocontador = 79
	ret.ventolimite = 15

	ret.ventofinalizador = ret.ventolimite

	return &ret
}

func (a *Ventos) ventoCorrenteNome() string {
	return a.VentosNome(a.vento)
}

func (a *Ventos) VentoCorrenteValor() float32 {
	return a.vento
}

func (a *Ventos) VentoCorrente() string {
	return a.VentosNome(a.vento)
}

func (a *Ventos) VentoOrigem() string {
	return a.ventoorigem
}

func (a *Ventos) VentoDestino() string {
	return a.ventodestino
}

func (a *Ventos) VentoRodando() bool {

	var ret bool = false
	if a.ventoorigem == a.ventodestino {
		ret = true
	}
	return ret
}

func (a *Ventos) VentoModo() string {

	var ventomodo string = ""

	if a.ventoorigem == a.ventodestino {
		ventomodo = "Brisa"

		if a.vento >= 0 && a.vento < 10 {
			ventomodo = "Brisa"
		}

		if a.vento >= 10 && a.vento < 20 {
			ventomodo = "Redimoinho"
		}

		if a.vento >= 20 && a.vento < 60 {
			ventomodo = "Ciclone"
		}

		if a.vento >= 60 && a.vento < 80 {
			ventomodo = "Tornado"
		}

		if a.vento >= 80 {
			ventomodo = "Furacão"
		}

	} else {

		ventomodo = " - "
	}

	return ventomodo
}

func (a *Ventos) VentosNome(_vento float32) string {
	var ret string = ""

	if _vento >= 0 && _vento < 10 {
		ret = "Sem Vento"
	}

	if _vento >= 10 && _vento < 20 {
		ret = "Leve"
	}

	if _vento >= 20 && _vento < 60 {
		ret = "Ventania"
	}

	if _vento >= 60 && _vento < 80 {
		ret = "Forte"
	}

	if _vento >= 80 {
		ret = "Fortissimo"
	}

	return ret
}

func (a *Ventos) Direcionar(mudardirecao int) string {
	var direcao string = ""

	if mudardirecao >= 0 && mudardirecao < 10 {
		direcao = "Norte"
	} else if mudardirecao >= 10 && mudardirecao < 20 {
		direcao = "Leste"
	} else if mudardirecao >= 10 && mudardirecao < 20 {
		direcao = "Sul"
	} else if mudardirecao >= 20 && mudardirecao < 30 {
		direcao = "Oeste"
	} else if mudardirecao >= 30 && mudardirecao < 40 {
		direcao = "Nordeste"
	} else if mudardirecao >= 40 && mudardirecao < 50 {
		direcao = "Sudeste"
	} else if mudardirecao >= 50 && mudardirecao < 60 {
		direcao = "Sudoeste"
	} else if mudardirecao >= 60 && mudardirecao < 70 {
		direcao = "Noroeste"
	}
	return direcao
}

func (a *Ventos) Ventar() {

	if a.ventocontador >= a.ventofinalizador {
		a.ventocontador = 0
		a.vento = float32(rand.Intn(int(100))) + rand.Float32()

		a.ventofinalizador = a.ventolimite + rand.Intn(30)

		//utils.Log("logs/ambientelimitador.txt", "Vento - "+strconv.Itoa(a.ventofinalizador)+"   "+a.ventoCorrenteNome())

		a.ventoorigem = a.Direcionar(rand.Intn(70))

		a.ventodestino = a.Direcionar(rand.Intn(70))

	} else {

		if rand.Intn(100) < 50 {
			a.vento += rand.Float32()
		} else {
			a.vento -= rand.Float32()
		}

	}

	if a.vento < 10 {
		a.ventoorigem = ""
		a.ventodestino = ""
	}
	a.ventocontador++
}

func (a *Ventos) VentoInfo() string {
	var ret string

	if a.ventoCorrenteNome() == "Sem Vento" {

		ret = fmt.Sprintf("Vento -  %.2f - Sem Vento", a.VentoCorrenteValor())

	} else {

		if a.VentoRodando() == true {
			ret = fmt.Sprintf("Vento -  %.2f ", a.VentoCorrenteValor()) + " : " + a.ventoCorrenteNome() + " [ " + a.ventoorigem + " -> " + a.ventodestino + " ] - " + a.VentoModo()
		} else {
			ret = fmt.Sprintf("Vento -  %.2f ", a.VentoCorrenteValor()) + " : " + a.ventoCorrenteNome() + " [ " + a.ventoorigem + " -> " + a.ventodestino + " ]"
		}

	}

	return ret
}
