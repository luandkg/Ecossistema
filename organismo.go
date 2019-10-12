package main

import (
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
}

func Organismonovo(nome string) *organismo {

	p := organismo{_nome: nome}
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"

	p._posx = 0
	p._posy = 0

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

func (p *organismo) atualizar(s *sdl.Surface) {

	var ni int32 = int32(p.x()) * 10
	var nj int32 = int32(p.y()) * 10
	p.rect = sdl.Rect{ni, nj, 10, 10}

	s.FillRect(&p.rect, 0xADFF2F)

}
