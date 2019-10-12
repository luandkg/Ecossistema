package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Projeto para Linguagens de Progrmacao 2019.02

// AUTOR : LUAN ALVES - 17/0003191
// AUTOR : MARUAN OLIVEIRA - 18/0057685

func main() {

	tb := Tabuleiro_novo("MATRIZ")

	tb.limpar()

	//lsplantas := list.New()
	var lsplantas [2]planta
	lsplantas[0] = *Planta_novo("Capim Gordura", 5)
	lsplantas[1] = *Planta_novo("Capim Verde", 10)

	// TODO: Criar forma generica de adicionar plantas e extrair em uma funcao ou metodo
	//lsplantas.PushBack(*Planta_novo("Capim Gordura", 5))
	//lsplantas.PushBack(*Planta_novo("Capim Verde", 10))

	// TODO: Extrair para metodo ou funcao o rand
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var x int = r1.Intn(50)
	var y int = r1.Intn(50)

	tb.mudar(x, y, 1)

	//produtores := &produtor{}
	//p1 := Produtor_Novo("Capim Gordura", 5)
	//produtores.Append(&p1)expression

	tb.mostrar()

	var ciclo int = 0

	for {

		fmt.Println("---------------- Ciclo :  ", ciclo, " --------------------------------")
		time.Sleep(time.Second)
		fmt.Println("")

		fmt.Println("PRODUTORES")

		for i := 0; i < 2; i++ {

			p := &lsplantas[i]

			fmt.Println("      - ", p.toString())

			if p.status() == "vivo" {
				p.vivendo()
			}

		}

		//for plantac := lsplantas.Front(); plantac != nil; plantac = plantac.Next() {
		//
		//	p := plantac.Value.(planta)
		//
		//	// TODO: adicionar toString para a struct
		//	fmt.Println("      - ", p.nome(), " [", p.fase(), ",", p.ciclos(), "]")
		//
		//	if p.status() == "vivo" {
		//		p.vivendo()
		//	}
		//
		//	//if capimverde.status() == "vivo" {
		//	//	capimverde.vivendo()
		//	//}
		//}

		ciclo++

		fmt.Println("")

		if ciclo >= 15 {
			break
		}

	}

	fmt.Println("Fim da Simulação !!!")

}
