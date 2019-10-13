package main

import (
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type animal struct {
	organismo
	_adultociclo int
}

func Animalnovo(nome string, adulto int) *animal {

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

func (p *animal) toString() string {

	var str = p.nome() + " [" + p.fase() + " " + strconv.Itoa(p.ciclos()) + "]" + " POS[" + strconv.Itoa(p.x()) + " " + strconv.Itoa(p.y()) + "]"

	return str
}

func (p *animal) atualizar(s *sdl.Surface) {

	var ni int32 = int32(p.x()) * 10
	var nj int32 = int32(p.y()) * 10
	p.rect = sdl.Rect{ni, nj, 10, 10}

	s.FillRect(&p.rect, 0xFF6347)

}
