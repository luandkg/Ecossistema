package main

import (
	"fmt"
)

type tabuleiro struct {

  _nome string
  dados [50][50]int

}

func Tabuleiro_novo(nome string) *tabuleiro {

  p := tabuleiro{_nome: nome}

	return &p
}


func (p *tabuleiro) limpar()  {

	// Zerando Tabuleiro
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			p.dados[i][j] = 0
		}
	}

}

func (p *tabuleiro) mostrar()  {

  fmt.Println("----------------------------------------- TABULEIRO ------------------------------------------------")

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {

			fmt.Print(" ", p.dados[i][j])
		}
		fmt.Print("\n")
	}

  fmt.Println("-----------------------------------------------------------------------------------------------------")


}

func (p *tabuleiro) mudar(x int,y int, valor int)  { p.dados[x][y] = valor}
