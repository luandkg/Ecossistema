package tabuleiro

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Tabuleiro struct {
	_nome string
	dados [50][50]peca
}

func TabuleiroNovo(nome string) *Tabuleiro {

	p := Tabuleiro{_nome: nome}

	var tamanhoTabuleiro = 50

	for i := 0; i < tamanhoTabuleiro; i++ {

		for j := 0; j < tamanhoTabuleiro; j ++ {

			p.dados[i][j]._x = i
			p.dados[i][j]._y = j

		}

	}

	return &p
}

func (p *Tabuleiro) Limpar() {

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

func (p *Tabuleiro) Atualizar(s *sdl.Surface) {

	// Zera surface rects
	s.FillRect(nil, 0)

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {

			p.dados[i][j].atualizar(s)

		}
	}

}

func (p *Tabuleiro) Mostrar() {

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

func (p *Tabuleiro) RecuperarPeca(x int, y int) *peca {

	return &p.dados[x][y]

}