package tabuleiro

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
		//s.FillRect(&p.rect, 0xFF4500)
	} else if p._valor == 1 {
		//fmt.Println("entrei no valor 1")
		//s.FillRect(&p.rect, 0x7CFC00)
	} else {
		//fmt.Println("entrei no valor else")
		//s.FillRect(&p.rect, 0x00BFFF)
	}

}

func (p *peca) VerificarPosicao() bool {

	if p._valor == 0 {
		return false
	} else {
		return true
	}

}

func (p *peca) OcuparPosicao() {
	p._valor = 1
}
