package ecossistema

import (
	"fmt"
	"math/rand"
	"tabuleiro"
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

func OrganismoNovo(nome string) *organismo {

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

func (p *organismo) Nome() string   { return p._nome }
func (p *organismo) Fase() string   { return p._fase }
func (p *organismo) Status() string { return p._status }
func (p *organismo) Ciclos() int    { return p._idade }

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

	var ni = int32(p.x()) * 10
	var nj = int32(p.y()) * 10
	p.rect = sdl.Rect{ni, nj, 10, 10}

	s.FillRect(&p.rect, p._cor)

}

func (p *organismo) Movimento(tb *tabuleiro.Tabuleiro) {

	p._dircontador++

	var tempX = p._posx
	var tempY = p._posy

	if p._dircontador >= p._dirquantidade {
		p._dircontador = 0
		p._direcao = ""
	}

	switch p._direcao {

	case "":
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		escolha := r1.Intn(3)
		p._dirquantidade = r1.Intn(15)

		switch escolha {

		case 0:
			p._direcao = "l"
			break

		case 1:
			p._direcao = "o"
			break

		case 2:
			p._direcao = "s"
			break

		case 3:
			p._direcao = "n"
			break

		}

		fmt.Println("Mudar direcao : ", p._direcao, "  com ", p._dirquantidade)

	case "l":
		tempX += 1
		if tempX >= 50 {
			p._direcao = "o"
			tempX = 48
		}
		break

	case "o":
		tempX -= 1
		if tempX < 0 {
			p._direcao = "l"
			tempX = 1
		}
		break

	case "n":
		tempY -= 1
		if tempY < 0 {
			p._direcao = "s"
			tempY = 1
		}
		break

	case "s":
		tempY += 1
		if tempY >= 50 {
			p._direcao = "n"
			tempY = 48
		}
		break

	}

	peca := tb.RecuperarPeca(tempX, tempY)

	if peca.VerificarPosicao() {

		p.Movimento(tb)

	} else {

		p._posx = tempX
		p._posy = tempY

	}

}
