package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Projeto para Linguagens de Programacao 2019.02

// AUTOR : LUAN ALVES - 17/0003191
// AUTOR : MARUAN OLIVEIRA - 18/0057685

var (
	window          *sdl.Window
	renderer        *sdl.Renderer
	event           sdl.Event
	err             error
	surface         *sdl.Surface

	running         bool
)

const (
	screenWidth  = 500
	screenHeight = 650
)

func Configuracao() (successful bool) {

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to initialize sdl: %s\n", err)
		return false
	}

	if err = ttf.Init(); err != nil {
		fmt.Printf("Failed to initialize TTF: %s\n", err)
		return false
	}

	window, err = sdl.CreateWindow("Ecossistema", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
		return false
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
		return false
	}

	surface, err = window.GetSurface()
	if err != nil {
		fmt.Printf("Failed to create surface: %s\n", err)
		return false
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")

	return true

}

func HandleEvents() {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				running = false
			}
		}
	}
}

func AtualizarTela(ecossistemaC *ecossistema) {

	RenderizarTextos(ecossistemaC)

	renderer.Present()

}

func Encerrar() {
	renderer.Destroy()
	window.Destroy()
	ttf.Quit()
	sdl.Quit()
}

func main() {

	if !Configuracao() {
		os.Exit(1)
	}

	// ESCOPO PRINCIPAL
	log("logs.txt", "")
	log("logs.txt", " ------------------ SIMULACAO ------------------ ")

	tb := Tabuleiro_novo("MATRIZ")
	ambienteC := AmbienteNovo()
	ecossistemaC := EcossistemaNovo()

	tb.limpar()

	// PLANTAS

	for i := 0; i < 10; i++ {
		ecossistemaC.adicionarProdutor(Plantanovo("Capim Gordura", 200, 100, 300, 0xADFF2F, ecossistemaC))
	}
	for i := 0; i < 10; i++ {
		ecossistemaC.adicionarProdutor(Plantanovo("Capim Verde", 300, 150, 600, 0x808000, ecossistemaC))
	}
	for i := 0; i < 10; i++ {
		ecossistemaC.adicionarProdutor(Plantanovo("Laranjeira", 500, 200, 10000, 0xDAA520, ecossistemaC))
	}
	for i := 0; i < 10; i++ {
		ecossistemaC.adicionarProdutor(Plantanovo("Ervacidreira", 300, 300, 1000, 0xFFFF00, ecossistemaC))
	}

	// ANIMAIS

	for i := 0; i < 10; i++ {
		ecossistemaC.adicionarConsumidor(Consumidor("Rato", 200, 200, 2000, 0xDDA0DD, ecossistemaC))
	}

	for i := 0; i < 4; i++ {
		ecossistemaC.adicionarConsumidor(Consumidor("Roeador", 400, 200, 5000, 0xEE82EE, ecossistemaC))
	}

	for i := 0; i < 6; i++ {
		ecossistemaC.adicionarConsumidor(Consumidor("Coelho", 500, 250, 8000, 0x7B68EE, ecossistemaC))
	}

	ecossistemaC.mapearOrganismos()

	running = true
	for running {

		HandleEvents()

		fmt.Println("---------------- Ciclo :  ", ambienteC.ciclo, " --------------------------------")
		time.Sleep(time.Second / 4)
		fmt.Println("")

		fmt.Println("PRODUTORES")

		tb.atualizar(surface, ambienteC)

		for p := range ecossistemaC.produtores {
			produtorc := ecossistemaC.produtores[p]

			if produtorc.status() == "vivo" {

				produtorc._nomecompleto = produtorc._nome + " " + p
				fmt.Println("      - ", produtorc.toString())
				produtorc.vivendo()
				produtorc.atualizar(surface)

			}

		}

		fmt.Println("CONSUMIDORES")

		for p := range ecossistemaC.consumidores {

			consumidorc := ecossistemaC.consumidores[p]

			if consumidorc.status() == "vivo" {

				fmt.Println("      - ", consumidorc.toString())
				consumidorc.vivendo()
				consumidorc.movimento()
				consumidorc.atualizar(surface)

			}

		}

		AtualizarTela(ecossistemaC)

		fmt.Println("")

		ambienteC.ambiente()

		fmt.Println("Fase -> ", ambienteC.fase)
		fmt.Println("Quantidade de Sol -> ", ambienteC.sol)
		fmt.Println("Ceu -> ", ambienteC.ceu())

		if ambienteC.fasecontador == 0 {

			log("logs.txt", "Plantas - "+strconv.Itoa(len(ecossistemaC.produtores)))
			log("logs.txt", "Consumidores - "+strconv.Itoa(len(ecossistemaC.consumidores)))

			ecossistemaC.removerOrganimosMortos()
		}
	}

	Encerrar()

	fmt.Println("Fim da Simulação !!!")

}
