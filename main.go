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
	lsplantas[0] = *Plantanovo("Capim Gordura", 5, 10, 16)
	lsplantas[1] = *Plantanovo("Capim Verde", 10, 20, 32)

	mapear(*tb, lsplantas)

	// TODO: Criar forma generica de adicionar plantas e extrair em uma funcao ou metodo
	//lsplantas.PushBack(*Planta_novo("Capim Gordura", 5))
	//lsplantas.PushBack(*Planta_novo("Capim Verde", 10))

	//produtores := &produtor{}
	//p1 := Produtor_Novo("Capim Gordura", 5)
	//produtores.Append(&p1)expression

	tb.mostrar()

	var ciclo int = 0
	var fase string = "Dia"
	var faseciclo int = 10
	var fasecontador int = 0

	for {

		fmt.Println("---------------- Ciclo :  ", ciclo, " --------------------------------")
		time.Sleep(time.Second)
		fmt.Println("")

		fmt.Println("PRODUTORES")

		for i := 0; i < 2; i++ {

			p := &lsplantas[i]

			if p.status() == "vivo" {

				fmt.Println("      - ", p.toString())
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

		if ciclo >= 50 {
			break
		}

		// Implementacao FASE - DIA / NOITE
		if fasecontador >= faseciclo {
			fasecontador = 0
			if fase == "Dia" {
				fase = "Noite"
			} else {
				fase = "Dia"
			}
		} else {
			fasecontador++
		}

		var sol int = 0
		var solmodo string = " - "
		if fase == "Dia" {
			r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
			sol = r1.Intn(100)

			if sol >= 0 && sol < 30 {
				solmodo = "Muito Nublado"
			}

			if sol >= 30 && sol < 50 {
				solmodo = "Nublado"
			}

			if sol >= 50 && sol < 70 {
				solmodo = "Normal"
			}

			if sol >= 70 && sol < 90 {
				solmodo = "Ensolado"
			}

			if sol >= 90 {
				solmodo = "Muito Ensolado"
			}

		} else {
			sol = 0
			solmodo = " - "
		}

		fmt.Println("Fase -> ", fase)
		fmt.Println("Quantidade de Sol -> ", sol)
		fmt.Println("Modo -> ", solmodo)

	}

	fmt.Println("Fim da Simulação !!!")

}

func mapear(tb tabuleiro, lsplantas [2]planta) {

	// Mapear plantas no Tabuleiro

	// TODO: Extrair para metodo ou funcao o rand
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 2; i++ {

		p := &lsplantas[i]

		var x int = r1.Intn(50)
		var y int = r1.Intn(50)

		p.mudarposicao(x, y)
		tb.mudar(x, y, 1)
		fmt.Println(" - ", x, " - ", y)
	}

}
