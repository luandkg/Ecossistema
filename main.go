package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Projeto para Linguagens de Progrmacao 2019.02

// AUTOR : LUAN ALVES - 17/0003191
// AUTOR : MARUAN OLIVEIRA - 18/0057685

var contadorplanta = 0
var contadoranimal = 0

func adicionarplanta(ls map[string]*planta, plantac *planta) {

	ls[strconv.Itoa(contadorplanta)] = plantac

	contadorplanta++
}

func adicionaranimal(ls map[string]*animal, animalc *animal) {

	ls[strconv.Itoa(contadoranimal)] = animalc

	contadoranimal++
}

func main() {

	// SDL INIT - Criacao da Janela
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Ecossistema", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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
	//var lsplantas [2]planta
	//lsplantas[0] = *Plantanovo("Capim Gordura", 5, 10, 16)
	//lsplantas[1] = *Plantanovo("Capim Verde", 10, 20, 32)
	//	mapear(tb, &lsplantas)

	//var lsanimais [2]animal
	//lsanimais[0] = *Animalnovo("Coelho", 5)
	//lsanimais[1] = *Animalnovo("Roedor", 10)

	//mapearanimais(tb, &lsanimais)

	// PLANTAS
	plantas := make(map[string]*planta)

	adicionarplanta(plantas, Plantanovo("Capim Gordura", 5, 10, 16))
	adicionarplanta(plantas, Plantanovo("Capim Verde", 10, 20, 32))
	adicionarplanta(plantas, Plantanovo("Laranjeira", 10, 20, 32))

	plantas["0"]._cor = 0xADFF2F
	plantas["1"]._cor = 0x808000
	plantas["2"]._cor = 0xDAA520

	for p := range plantas {

		var plantac = plantas[p]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		plantac.mudarposicao(x, y)
	}

	//var t = len(plantas)
	//fmt.Println("Tamanho : ", t)

	// ANIMAIS
	animais := make(map[string]*animal)

	adicionaranimal(animais, Animalnovo("Coelho", 50))
	adicionaranimal(animais, Animalnovo("Roeador", 30))
	adicionaranimal(animais, Animalnovo("Rato", 15))

	animais["0"]._cor = 0x7B68EE
	animais["1"]._cor = 0xEE82EE
	animais["2"]._cor = 0xDDA0DD

	for p := range animais {

		var animalc = animais[p]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		animalc.mudarposicao(x, y)
	}

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

		for p := range plantas {
			plantac := plantas[p]

			//if plantac.status() == "vivo" {

			fmt.Println("      - ", plantac.toString())
			plantac.vivendo()
			//p.movimento()
			plantac.atualizar(surface)
			window.UpdateSurface()

			//	}

		}

		fmt.Println("CONSUMIDORES")

		for p := range animais {

			animalc := animais[p]

			if animalc.status() == "vivo" {

				fmt.Println("      - ", animalc.toString())
				animalc.vivendo()
				animalc.movimento()
				animalc.atualizar(surface)
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
		fmt.Println("Ceu -> ", ceu())

	}

	fmt.Println("Fim da Simulação !!!")

}

func mapear(tb *tabuleiro, plantas map[string]planta) {

	// Mapear plantas no Tabuleiro

	for p := range plantas {

		plantac := plantas[p]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		plantac.mudarposicao(x, y)
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
