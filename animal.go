package main

import (
	"fmt"
)

type animal struct {
	organismo
	_adultociclo int
}

func Animal_novo(nome string, adulto int) *animal {

	p := animal{_adultociclo: adulto}
	p._nome = nome
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

	return &p
}

func (p *animal) vivendo() {

	p.organismo.vivendo()

	if p._status == "vivo" {

		if p._fase == "nascido" {
			if p._idade >= p._adultociclo {
				p._fase = "adulto"

				fmt.Println("       --- Produtor : ", p.nome(), " Evoluiu : Adulto !!!")

			}
		}

	}

}
