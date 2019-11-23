package ecossistema

import (
	"fmt"
	"math/rand"
	"utils"
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

		if a.vento >= 10 && a.vento < 20 {
			ventomodo = "Ciclone"
		}

		if a.vento >= 20 && a.vento < 60 {
			ventomodo = "Anticiclone"
		}

		if a.vento >= 60 && a.vento < 80 {
			ventomodo = "Mesocilone"
		}

		if a.vento >= 80 {
			ventomodo = "Tornado"
		}

	} else {

		ventomodo = " - "
	}

	return ventomodo
}

func (a *Ventos) VentosNome(_vento float32) string {
	var ret string = ""

	// A Escala de Beaufort classifica a intensidade dos ventos, tendo em conta a sua velocidade e os efeitos resultantes das ventanias no mar e em terra.
	// Foi concebida pelo meteorologista anglo-irlandês Francis Beaufort no início do século XIX.
	// Na década de 1830, a escala de Beaufort já era amplamente utilizada pela Marinha Real Britânica.

	if _vento < 10 {
		ret = "Sem Vento"
	}

	if _vento >= 10 && _vento < 15 {
		ret = "Calmo"
	}

	if _vento >= 15 && _vento < 20 {
		ret = "Aragem"
	}

	if _vento >= 20 && _vento < 25 {
		ret = "Brisa Leve"
	}

	if _vento >= 25 && _vento < 30 {
		ret = "Brisa Fraca"
	}

	if _vento >= 30 && _vento < 35 {
		ret = "Brisa Moderada"
	}

	if _vento >= 35 && _vento < 40 {
		ret = "Brisa Forte"
	}

	if _vento >= 40 && _vento < 45 {
		ret = "Vento Fresco"
	}

	if _vento >= 45 && _vento < 50 {
		ret = "Vento Forte"
	}

	if _vento >= 50 && _vento < 60 {
		ret = "Ventania"
	}

	if _vento >= 60 && _vento < 70 {
		ret = "Ventania Forte"
	}

	if _vento >= 70 && _vento < 80 {
		ret = "Tempestade"
	}

	if _vento >= 80 && _vento < 90 {
		ret = "Tempestade Violenta"
	}

	if _vento >= 90 {
		ret = "Furacão"
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

func (a *Ventos) VentoVelocidadeCorrenteInfo() string {
	return fmt.Sprintf("%.2f Km/h", a.VentoVelocidadeCorrente())
}

func (a *Ventos) VentoVelocidadeCorrente() float32 {
	return a.VentoVelocidade(a.vento)
}

func (a *Ventos) VentoVelocidade(_vento float32) float32 {
	var ret float32 = 0

	if _vento < 10 {
		ret = 0
	}

	if _vento >= 10 && _vento < 15 {
		ret = rand.Float32()
	}

	if _vento >= 15 && _vento < 20 {
		ret = float32(1+utils.Aleatorionumero(4)) + rand.Float32()
	}

	if _vento >= 20 && _vento < 25 {
		ret = float32(5+utils.Aleatorionumero(6)) + rand.Float32()
	}

	if _vento >= 25 && _vento < 30 {
		ret = float32(11+utils.Aleatorionumero(8)) + rand.Float32()
	}

	if _vento >= 30 && _vento < 35 {
		ret = float32(19+utils.Aleatorionumero(2)) + rand.Float32()
	}

	if _vento >= 35 && _vento < 40 {
		ret = float32(28+utils.Aleatorionumero(4)) + rand.Float32()
	}

	if _vento >= 40 && _vento < 45 {
		ret = float32(38+utils.Aleatorionumero(6)) + rand.Float32()
	}

	if _vento >= 45 && _vento < 50 {
		ret = float32(49+utils.Aleatorionumero(8)) + rand.Float32()
	}

	if _vento >= 50 && _vento < 60 {
		ret = float32(61+utils.Aleatorionumero(2)) + rand.Float32()
	}

	if _vento >= 60 && _vento < 70 {
		ret = float32(74+utils.Aleatorionumero(4)) + rand.Float32()
	}

	if _vento >= 70 && _vento < 80 {
		ret = float32(88+utils.Aleatorionumero(6)) + rand.Float32()
	}

	if _vento >= 80 && _vento < 90 {
		ret = float32(102+utils.Aleatorionumero(8)) + rand.Float32()
	}

	if _vento >= 90 {
		ret = float32(117+utils.Aleatorionumero(10)) + rand.Float32()
	}

	return ret
}
