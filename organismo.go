package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type organismo struct {
	_nome   string
	_idade  int
	_status string
	_fase   string

	_posx int
	_posy int

	rect sdl.Rect

	_direcao       string
	_dirquantidade int
	_dircontador   int

	_cor uint32
}

func Organismonovo(nome string) *organismo {

	p := organismo{_nome: nome}
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"

	p._posx = 0
	p._posy = 0

	p._cor = 0xADFF2F

	p.rect = sdl.Rect{0, 0, 10, 10}

	return &p
}

func (p *organismo) vivendo() {

	if p._status == "vivo" {

		p._idade++

		if p._fase == "nascido" {

		}

	}

}

func (p *organismo) nome() string   { return p._nome }
func (p *organismo) fase() string   { return p._fase }
func (p *organismo) status() string { return p._status }
func (p *organismo) ciclos() int    { return p._idade }

func (p *organismo) mudarposicao(x int, y int) {
	p._posx = x
	p._posy = y

}

func (p *organismo) x() int { return p._posx }
func (p *organismo) y() int { return p._posy }

func (p *organismo) setCor(cor uint32) {
	p._cor = cor
}

func (p *organismo) atualizar(s *sdl.Surface) {

	var ni int32 = int32(p.x()) * 10
	var nj int32 = int32(p.y()) * 10
	p.rect = sdl.Rect{ni, nj, 10, 10}

	s.FillRect(&p.rect, p._cor)

}

func (p *organismo) movimento() {

	p._dircontador++

	if p._dircontador >= p._dirquantidade {
		p._dircontador = 0
		p._direcao = ""
	}

	if p._direcao == "" {
		p._direcao = "l"

		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		escolha := r1.Intn(3)
		p._dirquantidade = r1.Intn(15)

		if escolha == 0 {
			p._direcao = "l"
		}

		if escolha == 1 {
			p._direcao = "o"
		}

		if escolha == 2 {
			p._direcao = "s"
		}

		if escolha == 3 {
			p._direcao = "n"
		}

		fmt.Println("Mudar direcao : ", p._direcao, "  com ", p._dirquantidade)
	}

	if p._direcao == "l" {
		p._posx += 1
		if p._posx >= 50 {
			p._direcao = "o"
			p._posx = 48
		}
	} else if p._direcao == "o" {
		p._posx -= 1

		if p._posx < 0 {
			p._direcao = "l"
			p._posx = 1
		}

	} else if p._direcao == "n" {
		p._posy -= 1

		if p._posy < 0 {
			p._direcao = "s"
			p._posy = 1
		}
	} else if p._direcao == "s" {
		p._posy += 1

		if p._posy >= 50 {
			p._direcao = "n"
			p._posy = 48
		}

	}
}
