package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type peca struct {
	_x     int
	_y     int
	rect   sdl.Rect
	_valor int
}

func (p *peca) atualizar(s *sdl.Surface) {

	if p._valor == 0 {
		//	s.FillRect(&p.rect, 0xFF4500)
	} else if p._valor == 1 {
		s.FillRect(&p.rect, 0x7CFC00)

	} else {
		s.FillRect(&p.rect, 0x00BFFF)

	}

}
