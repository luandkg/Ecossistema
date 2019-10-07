package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Projeto para Linguagens de Progrmacao 2019.02

// AUTOR : LUAN ALVES - 17/0003191
// AUTOR : MARUAN -

func main() {

	tb := Tabuleiro_novo("MATRIZ")

	tb.limpar()

	capimgorgura := Produtor_Novo("Capim Gordura", 5)
	capimverde := Produtor_Novo("Capim Verde", 10)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var x int = r1.Intn(50)
	var y int = r1.Intn(50)

	tb.mudar(x, y, 1)

	//produtores := &produtor{}
	//p1 := Produtor_Novo("Capim Gordura", 5)
	//produtores.Append(&p1)

	tb.mostrar()

	var ciclo int = 0

	for {

		fmt.Println("---------------- Ciclo :  ", ciclo, " --------------------------------")
		time.Sleep(time.Second)
		fmt.Println("")

		fmt.Println("PRODUTORES")
		fmt.Println("      - ", capimgorgura.nome(), " [", capimgorgura.fase(), ",", capimgorgura.ciclos(), "]")
		fmt.Println("      - ", capimverde.nome(), " [", capimverde.fase(), ",", capimverde.ciclos(), "]")

		if capimgorgura.status() == "vivo" {
			capimgorgura.vivendo()
		}

		if capimverde.status() == "vivo" {
			capimverde.vivendo()
		}

		ciclo++

		fmt.Println("")

		if ciclo >= 15 {
			break
		}

	}

	fmt.Println("Fim da Simulação !!!")

}
