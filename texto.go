package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	textosTextures  []texto
)

type texto struct {
	largura int32
	altura int32
	textura *sdl.Texture
}

func CriarTextosTexturas(textos []string) (successful bool) {

	textosTextures = nil

	var font *ttf.Font

	if font, err = ttf.OpenFont("./assets/fonts/OpenSans-Regular.ttf", 14); err != nil {
		fmt.Printf("Failed to open font: %s\n", err)
		return false
	}

	for i := 0; i < len(textos); i++ {

		var solidSurface *sdl.Surface
		if solidSurface, err = font.RenderUTF8Solid(textos[i], sdl.Color{255, 0, 0, 255}); err != nil {
			fmt.Printf("Failed to render text: %s\n", err)
			return false
		}

		textoLargura, textoAltura, err := font.SizeUTF8(textos[i])

		fmt.Println("----------------- FONT -------------------------")
		fmt.Println(textoLargura, textoAltura)

		var novoTextoTextura *sdl.Texture

		if novoTextoTextura, err = renderer.CreateTextureFromSurface(solidSurface); err != nil {
			fmt.Printf("Failed to create texture: %s\n", err)
			return false
		}

		var textoRenderizado = texto{int32(textoLargura), int32(textoAltura), novoTextoTextura}

		textosTextures = append(textosTextures, textoRenderizado)

		solidSurface.Free()

	}

	font.Close()

	return true

}

func RenderizarTextos(ecossistemaC *ecossistema) {

	var produtoresTotal = fmt.Sprintf("Produtores: %d", len(ecossistemaC.produtores))
	var consumidoresTotal = fmt.Sprintf("Consumidores: %d", len(ecossistemaC.consumidores))

	var testando = []string{produtoresTotal, consumidoresTotal}

	if !CriarTextosTexturas(testando) {
		os.Exit(2)
	}

	for i := 0; i < len(textosTextures); i++ {

		var x = int32(0)
		var y = int32(550 + i * 30)
		var largura = textosTextures[i].largura
		var altura = textosTextures[i].altura

		renderer.Copy(textosTextures[i].textura, nil, &sdl.Rect{x, y, largura, altura})

	}

}