package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type tabuleiro struct {
	_nome string
	dados [50][50]peca
}

func Tabuleiro_novo(nome string) *tabuleiro {

	p := tabuleiro{_nome: nome}

	return &p
}

func (p *tabuleiro) limpar() {

	// Zerando Tabuleiro
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {

			var ni int32 = int32(i) * 10
			var nj int32 = int32(j) * 10
			p.dados[i][j].rect = sdl.Rect{ni, nj, 10, 10}
			p.dados[i][j]._valor = 0
		}
	}

}

func (p *tabuleiro) atualizar(s *sdl.Surface) {

	s.FillRect(nil, 0)

	// Zerando Tabuleiro
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {

			p.dados[i][j].atualizar(s)

		}
	}

	s.FillRect(&p.dados[0][0].rect, 0x00BFFF)
	//s.FillRect(&p.dados[49][0].rect, 0x00BFFF)

	//	s.FillRect(&p.dados[0][49].rect, 0x00BFFF)
	//0xFFA500s.FillRect(&p.dados[49][49].rect, 0x00BFFF)

	var linhafinal = sdl.Rect{0, 500, 500, 10}
	if fase == "Dia" {
		s.FillRect(&linhafinal, 0xFFFF00)
	} else {
		s.FillRect(&linhafinal, 0x000080)
	}

	var st int = sol * 5
	var solinha = sdl.Rect{0, 510, int32(st), 10}
	s.FillRect(&solinha, 0xFF4500)

}

func (p *tabuleiro) mostrar() {

	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("\n")

	fmt.Println("----------------------------------------- TABULEIRO ------------------------------------------------")
	fmt.Println("\n")
	fmt.Println("\n")

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {

			fmt.Print(" ", p.dados[i][j]._valor)
		}
		fmt.Print("\n")
	}

	fmt.Println("\n")
	fmt.Println("\n")

	fmt.Println("-----------------------------------------------------------------------------------------------------")

}

func (p *tabuleiro) mudar(x int, y int, valor int) {
	//p.dados[x][y]._valor = valor
}
