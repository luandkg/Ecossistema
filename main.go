package main

import (
	"fmt"
	"math/rand"
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
		500, 500, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, 0)

	// ESCOPO PRINCIPAL0xFFA500
	log("logs.txt", "")
	log("logs.txt", " ------------------ SIMULACAO ------------------ ")

	tb := Tabuleiro_novo("MATRIZ")

	tb.limpar()

	//lsplantas := list.New()
	var lsplantas [2]planta
	lsplantas[0] = *Plantanovo("Capim Gordura", 5, 10, 16)
	lsplantas[1] = *Plantanovo("Capim Verde", 10, 20, 32)

	mapear(tb, &lsplantas)

	// TODO: Criar forma generica de adicionar plantas e extrair em uma funcao ou metodo
	//lsplantas.PushBack(*Planta_novo("Capim Gordura", 5))
	//lsplantas.PushBack(*Planta_novo("Capim Verde", 10))

	//produtores := &produtor{}
	//p1 := Produtor_Novo("Capim Gordura", 5)
	//produtores.Append(&p1)expression

	tb.mostrar()

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
		time.Sleep(time.Second)
		fmt.Println("")

		fmt.Println("PRODUTORES")

		tb.atualizar(surface)
		window.UpdateSurface()

		for i := 0; i < 2; i++ {

			p := &lsplantas[i]

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

		//}if ciclo >= 60 {
		//	break
		//}

		ambiente()

		fmt.Println("Fase -> ", fase)
		fmt.Println("Quantidade de Sol -> ", sol)
		fmt.Println("Modo -> ", solmodo)

	}

	fmt.Println("Fim da Simulação !!!")

}

func mapear(tb *tabuleiro, lsplantas *[2]planta) {

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
