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

var contadorplanta = 0
var contadoranimal = 0

var (
	window          *sdl.Window
	renderer        *sdl.Renderer
	event           sdl.Event
	err             error
	solidTexture    *sdl.Texture
	blendedTexture  *sdl.Texture
	shadedTexture   *sdl.Texture
	surface         *sdl.Surface

	running         bool
)

const (
	screenWidth  = 500
	screenHeight = 650
)

func adicionarplanta(ls map[string]*planta, plantac *planta) {

	ls[strconv.Itoa(contadorplanta)] = plantac

	contadorplanta++
}

func adicionaranimal(ls map[string]*consumidor, animalc *consumidor) {

	ls[strconv.Itoa(contadoranimal)] = animalc

	contadoranimal++
}

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

func CriarTextos() (successful bool) {
	var font *ttf.Font

	if font, err = ttf.OpenFont("./assets/fonts/OpenSans-Regular.ttf", 40); err != nil {
		fmt.Printf("Failed to open font: %s\n", err)
		return false
	}

	var solidSurface *sdl.Surface
	if solidSurface, err = font.RenderUTF8Solid("Deu certo!!", sdl.Color{255, 0, 0, 255}); err != nil {
		fmt.Printf("Failed to render text: %s\n", err)
		return false
	}

	if solidTexture, err = renderer.CreateTextureFromSurface(solidSurface); err != nil {
		fmt.Printf("Failed to create texture: %s\n", err)
		return false
	}

	solidSurface.Free()

	font.Close()

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

func AtualizarTela() {

	//X, Y, width, height
	renderer.Copy(solidTexture, nil, &sdl.Rect{0, 550, 190, 53})

	renderer.Present()

}

func Encerrar() {
	solidTexture.Destroy()
	shadedTexture.Destroy()
	blendedTexture.Destroy()
	renderer.Destroy()
	window.Destroy()
	ttf.Quit()
	sdl.Quit()
}

func main() {

	if !Configuracao() {
		os.Exit(1)
	}

	if !CriarTextos() {
		os.Exit(2)
	}

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

	// ANIMAIS
	animais := make(map[string]*consumidor)

	for i := 0; i < 10; i++ {
		adicionaranimal(animais, Consumidor("Rato", 15, 10, 30, 0xDDA0DD, animais))
	}

	for i := 0; i < 4; i++ {
		adicionaranimal(animais, Consumidor("Roeador", 30, 10, 30, 0xEE82EE, animais))
	}

	for i := 0; i < 6; i++ {
		adicionaranimal(animais, Consumidor("Coelho", 50, 10, 30,0x7B68EE, animais))
	}

	for p := range animais {

		var animalc = animais[p]

		var x = aleatorionumero(50)
		var y = aleatorionumero(50)

		animalc.mudarposicao(x, y)
	}

	var ciclo int = 0

	running = true
	for running {

		HandleEvents()

		fmt.Println("---------------- Ciclo :  ", ciclo, " --------------------------------")
		time.Sleep(time.Second / 4)
		fmt.Println("")

		fmt.Println("PRODUTORES")

		tb.atualizar(surface)

		for p := range plantas {
			plantac := plantas[p]

			if plantac.status() == "vivo" {

				fmt.Println("      - ", plantac.toString())
				plantac.vivendo()
				plantac.atualizar(surface)

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

			}

		}

		ciclo++

		AtualizarTela()

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

	Encerrar()

	fmt.Println("Fim da Simulação !!!")

}
