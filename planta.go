package main

import (
	"fmt"
)

type planta struct {
	organismo
	_adultociclo        int
	_reproduzirciclo    int
	_reproduzircontador int
}

// Plantanovo : Criar instancia
func Plantanovo(nome string, adulto int, reproducao int) *planta {

	p := planta{_adultociclo: adulto}
	p._nome = nome
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

	p._reproduzirciclo = reproducao
	p._reproduzircontador = 0

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

		if p._fase == "adulto" {

			p._reproduzircontador += 1

			if p._reproduzircontador >= p._reproduzirciclo {
				p._reproduzircontador = 0
				fmt.Println("       --- Planta : ", p.nome(), " Reproduzindo !!!")
			}

		}

	}

}

func (p *planta) toString() string {

	return fmt.Sprintf("%s [%s, %d]", p.nome(), p.fase(), p.ciclos())

}
