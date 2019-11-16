package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	nomeProjeto   = "Ecossistema"
	janelaLargura = 920
	janelaAltura  = 800
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

	window, err = sdl.CreateWindow(nomeProjeto, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, janelaLargura, janelaAltura, sdl.WINDOW_SHOWN)
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

func ManipularEventos() {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			running = false
		}
	}
}

func Encerrar() {
	renderer.Destroy()
	window.Destroy()
	font.Close()
	ttf.Quit()
	sdl.Quit()
}
