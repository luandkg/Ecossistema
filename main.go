package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Projeto para Linguagens de Progrmacao 2019.02

// AUTOR : LUAN ALVES - 17/0003191
// AUTOR : MARUAN OLIVEIRA - 18/0057685

func main() {

	// SDL INIT - Criacao da Janela
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Evolucao", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		500, 550, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, 0)

	// ESCOPO PRINCIPAL
	log("logs.txt", "")
	log("logs.txt", " ------------------ SIMULACAO ------------------ ")

	tb := Tabuleiro_novo("MATRIZ")

	tb.limpar()

	//lsplantas := list.New()
	var lsplantas [2]planta
	lsplantas[0] = *Plantanovo("Capim Gordura", 5, 10, 16)
	lsplantas[1] = *Plantanovo("Capim Verde", 10, 20, 32)

	mapear(tb, &lsplantas)

	var lsanimais [2]animal
	lsanimais[0] = *Animalnovo("Coelho", 5)
	lsanimais[1] = *Animalnovo("Roedor", 10)

	mapearanimais(tb, &lsanimais)

	// TODO: Criar forma generica de adicionar plantas e extrair em uma funcao ou metodo
	//lsplantas.PushBack(*Planta_novo("Capim Gordura", 5))
	//lsplantas.PushBack(*Planta_novo("Capim Verde", 10))

	//produtores := &produtor{}
	//p1 := Produtor_Novo("Capim Gordura", 5)
	//produtores.Append(&p1)expression

	//tb.mostrar()

	var ciclo int = 0

	running := true
	for running {

		// TRATAMENTO DE EVENTOS
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		fmt.Println("---------------- Ciclo :  ", ciclo, " --------------------------------")
		time.Sleep(time.Second / 4)
		fmt.Println("")

		fmt.Println("PRODUTORES")

		tb.atualizar(surface)
		window.UpdateSurface()

		for i := 0; i < 2; i++ {

			p := &lsplantas[i]

			if p.status() == "vivo" {

				fmt.Println("      - ", p.toString())
				p.vivendo()
				//p.movimento()
				p.atualizar(surface)
				window.UpdateSurface()

			}

		}

		for i := 0; i < 2; i++ {

			p := &lsanimais[i]

			if p.status() == "vivo" {

				fmt.Println("      - ", p.toString())
				p.vivendo()
				p.movimento()
				p.atualizar(surface)
				window.UpdateSurface()

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

		ambiente()

		fmt.Println("Fase -> ", fase)
		fmt.Println("Quantidade de Sol -> ", sol)
		fmt.Println("Modo -> ", solmodo)

	}

	fmt.Println("Fim da Simulação !!!")

}

func mapear(tb *tabuleiro, lsplantas *[2]planta) {

	// Mapear plantas no Tabuleiro

	for i := 0; i < 2; i++ {

		p := &lsplantas[i]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		p.mudarposicao(x, y)

	}

}

func mapearanimais(tb *tabuleiro, ls *[2]animal) {

	// Mapear plantas no Tabuleiro

	for i := 0; i < 2; i++ {

		p := &ls[i]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		p.mudarposicao(x, y)

	}

}
