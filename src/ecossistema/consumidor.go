package ecossistema

import (
	"fmt"
	"strconv"

	"utils"
)

type Consumidor struct {
	organismo
	_adultociclo        int
	_reproduzirciclo    int
	_reproduzircontador int

	_vida int

	_ecossistemaC *Ecossistema
}

func ConsumidorNovo(nome string, adulto int, reproducao int, vida int, cor uint32, ecossistemaC *Ecossistema) *Consumidor {

	p := Consumidor{_adultociclo: adulto}

	p.organismo = *OrganismoNovo(nome)
	p._nome = nome
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

	p._reproduzirciclo = reproducao
	p._reproduzircontador = 0

	p._vida = vida
	p._posx = 0
	p._posy = 0

	p._cor = cor
	p._ecossistemaC = ecossistemaC

	return &p

}

func (p *Consumidor) vivendo() {

	p.organismo.vivendo()

	if p._status == "vivo" {

		if p._idade == p._adultociclo || p._idade == p._vida {

			p.mudarFase()

		}

		if p._fase == "adulto" && p._idade < p._vida {

			p.reproduzir()

		}

		if p._idade >= p._vida {
			p._status = "morto"
			fmt.Println("       --- Consumidor : ", p.Nome(), " Morreu !!!")
		}
	}

}

func (p *Consumidor) mudarFase() {

	switch p._fase {

	case "nascido":
		p._fase = "adulto"
		fmt.Println("       --- Consumidor : ", p.Nome(), " Evoluiu : Adulto !!!")
		break

	case "adulto":
		p._status = "morto"
		p._fase = "falescido"
		fmt.Println("       --- Consumidor : ", p.Nome(), " Morreu !!!")
		break

	}

}

func (p *Consumidor) reproduzir() {

	p._reproduzircontador += 1

	if p._reproduzircontador >= p._reproduzirciclo {
		p._reproduzircontador = 0
		fmt.Println("       --- Consumidor : ", p.Nome(), " Reproduzindo !!!")

		var pg = ConsumidorNovo("Coelho", 5, 10, 16, 0xADFF2F, p._ecossistemaC)
		var x int = utils.Aleatorionumero(50)
		var y int = utils.Aleatorionumero(50)

		pg.mudarposicao(x, y)

		p._ecossistemaC.AdicionarConsumidor(pg)
	}

}

func (p *Consumidor) toString() string {

	var str = p.Nome() + " [" + p.Fase() + " " + strconv.Itoa(p.Ciclos()) + "]" + " POS[" + strconv.Itoa(p.x()) + " " + strconv.Itoa(p.y()) + "]"

	return str
}
