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

	// PLANTAS
	plantas := make(map[string]*planta)

	for i := 0; i < 10; i++ {
		adicionarplanta(plantas, Plantanovo("Capim Gordura", 200, 100, 300, 0xADFF2F, plantas))
	}
	for i := 0; i < 10; i++ {
		adicionarplanta(plantas, Plantanovo("Capim Verde", 300, 150, 600, 0x808000, plantas))
	}
	for i := 0; i < 10; i++ {
		adicionarplanta(plantas, Plantanovo("Laranjeira", 500, 200, 10000, 0xDAA520, plantas))
	}
	for i := 0; i < 10; i++ {
		adicionarplanta(plantas, Plantanovo("Ervacidreira", 300, 300, 1000, 0xFFFF00, plantas))
	}

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

	for i := 0; i < 10; i++ {
		adicionaranimal(animais, Animalnovo("Rato", 15, 0xDDA0DD))
	}

	for i := 0; i < 4; i++ {
		adicionaranimal(animais, Animalnovo("Roeador", 30, 0xEE82EE))
	}

	for i := 0; i < 6; i++ {
		adicionaranimal(animais, Animalnovo("Coelho", 50, 0x7B68EE))
	}

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

			if plantac.status() == "vivo" {

				fmt.Println("      - ", plantac.toString())
				plantac.vivendo()
				//p.movimento()
				plantac.atualizar(surface)
				window.UpdateSurface()

			}

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

		ciclo++

		fmt.Println("")

		ambiente()

		fmt.Println("Fase -> ", fase)
		fmt.Println("Quantidade de Sol -> ", sol)
		fmt.Println("Ceu -> ", ceu())

		if fasecontador == 0 {

			log("logs.txt", "Plantas - "+strconv.Itoa(len(plantas)))
			log("logs.txt", "Animais - "+strconv.Itoa(len(animais)))

		}
	}

	fmt.Println("Fim da Simulação !!!")

}
