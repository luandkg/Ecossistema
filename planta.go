package main

import (
	"fmt"
	"strconv"
)

type planta struct {
	organismo
	_adultociclo        int
	_reproduzirciclo    int
	_reproduzircontador int

	_vida int

	_plantas (map[string]*planta)
}

// Plantanovo : Criar instancia de planta
func Plantanovo(nome string, adulto int, reproducao int, vida int, cor uint32, plantas map[string]*planta) *planta {

	p := planta{_adultociclo: adulto}

	p.organismo = *Organismonovo(nome)
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
	p._plantas = plantas

	return &p
}

func (p *planta) vivendo(ecossistemaC *ecossistema) {

	p.organismo.vivendo()

	if p._status == "vivo" {

		if p._fase == "nascido" {
			if p._idade >= p._adultociclo {
				p._fase = "adulto"

				fmt.Println("       --- Planta : ", p.nome(), " Evoluiu : Adulto !!!")

			}
		}

		// Se o organismo for adulto inicia o ciclo de reproducao
		if p._fase == "adulto" {

			p._reproduzircontador += 1

			if p._reproduzircontador >= p._reproduzirciclo {
				p._reproduzircontador = 0
				fmt.Println("       --- Planta : ", p.nome(), " Reproduzindo !!!")

				var pg = Plantanovo(p._nome, p._adultociclo, p._reproduzirciclo, p._vida, p._cor, p._plantas)
				var x int = aleatorionumero(50)
				var y int = aleatorionumero(50)

				pg.mudarposicao(x, y)

				ecossistemaC.adicionarPlanta(pg)
			}

		}

		if p._idade >= p._vida {
			p._status = "morto"
			fmt.Println("       --- Planta : ", p.nome(), " Morreu !!!")
		}

	}

}

func (p *planta) toString() string {

	var str = p.nome() + " [" + p.fase() + " " + strconv.Itoa(p.ciclos()) + "]" + " POS[" + strconv.Itoa(p.x()) + " " + strconv.Itoa(p.y()) + "]"

	return str
}
