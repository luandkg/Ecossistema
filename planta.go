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
}

// Plantanovo : Criar instancia de planta
func Plantanovo(nome string, adulto int, reproducao int, vida int) *planta {

	p := planta{_adultociclo: adulto}
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

	return &p
}

func (p *planta) vivendo() {

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
			}

		}

		if p._idade >= p._vida {
			//p._status = "morto"
			fmt.Println("       --- Planta : ", p.nome(), " Morreu !!!")
		}

	}

}

func (p *planta) toString() string {

	var str = p.nome() + " [" + p.fase() + " " + strconv.Itoa(p.ciclos()) + "]" + " POS[" + strconv.Itoa(p.x()) + " " + strconv.Itoa(p.y()) + "]"

	return str
}
