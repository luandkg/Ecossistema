package main

import (
	"fmt"
)

type planta struct {
	organismo
	_adultociclo int
}

func Planta_novo(nome string, adulto int) *planta {

	p := planta{_adultociclo: adulto}
	p._nome = nome
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

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

	}

}
