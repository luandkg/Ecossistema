package main

import (
	"ecossistema"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
)

var (
	textosTextures  []texto
	textoLargura    int
	textoAltura     int
)

type texto struct {
	largura int32
	altura int32
	textura *sdl.Texture
}

func criarTextosTexturas(textos []string) (successful bool) {

	dir, err := os.Getwd()

	//dir=filepath.Dir(dir)
	//dir=filepath.Dir(dir)


	//fmt.Println("Local da Fonte : " + dir +"/assets/fonts/OpenSans-Regular.ttf")
	if err != nil {
		fmt.Printf("Failed to open get current directory: %s\n", err)
		return false
	}

	textosTextures = nil

	if font, err = ttf.OpenFont( dir +"/assets/fonts/OpenSans-Regular.ttf", 14); err != nil {
		fmt.Printf("Failed to open font: %s\n", err)
		return false
	}

	for i := 0; i < len(textos); i++ {

		var solidSurface *sdl.Surface
		if solidSurface, err = font.RenderUTF8Solid(textos[i], sdl.Color{255, 255, 255, 255}); err != nil {
			fmt.Printf("Failed to render text: %s\n", err)
			return false
		}

		var novoTextoTextura *sdl.Texture
		if novoTextoTextura, err = renderer.CreateTextureFromSurface(solidSurface); err != nil {
			fmt.Printf("Failed to create texture: %s\n", err)
			return false
		}

		if textoLargura, textoAltura, err = font.SizeUTF8(textos[i]); err != nil {
			fmt.Printf("Failed to get width or height: %s\n", err)
			return false
		}

		var textoRenderizado = texto{int32(textoLargura), int32(textoAltura), novoTextoTextura}

		textosTextures = append(textosTextures, textoRenderizado)

		solidSurface.Free()

	}

	return true

}

func carregarTextos(e *ecossistema.Ecossistema) []string {

	totalProdutoresJovens, totalProdutoresAdulto := e.TotalProdutoresFase()
	var produtoresTotal = fmt.Sprintf("Produtores: %d", e.TotalProdutores())
	var produtoresJovens = fmt.Sprintf("Prod. Jovens: %d", totalProdutoresJovens)
	var produtoresAdultos = fmt.Sprintf("Prod. Adultos: %d", totalProdutoresAdulto)

	totalConsumidoresJovens, totalConsumidoresAdulto := e.TotalConsumidoresFase()
	var consumidoresTotal = fmt.Sprintf("Consumidores: %d", e.TotalConsumidores())
	var consumidoresJovens = fmt.Sprintf("Cons. Jovens: %d", totalConsumidoresJovens)
	var consumidoresAdultos = fmt.Sprintf("Cons. Adultos: %d", totalConsumidoresAdulto)

	return []string{produtoresTotal, produtoresJovens, produtoresAdultos, consumidoresTotal, consumidoresJovens, consumidoresAdultos}

}

func RenderizarTextos(e *ecossistema.Ecossistema) {

	var textosParaRenderizar = carregarTextos(e)

	if !criarTextosTexturas(textosParaRenderizar) {
		os.Exit(2)
	}

	var maxColunas = 3
	var inicioY = 560
	var alturaMaximaY = 20
	var espacoY = 20

	var inicioX = 10
	var espacoX = inicioX
	var larguraMaximaX = 470 / maxColunas

	var colunaAtual = 1

	for i := 0; i < len(textosTextures); i++ {

		if i % maxColunas == 0 && i != 0 {
			espacoX = inicioX
			espacoY += alturaMaximaY
			colunaAtual = 1
		} else {
			if i != 0 {
				espacoX = larguraMaximaX * colunaAtual
				colunaAtual += colunaAtual
			}
		}

		var x = int32(inicioX + espacoX)
		var y = int32(inicioY + espacoY)
		var largura = textosTextures[i].largura
		var altura = textosTextures[i].altura

		renderer.Copy(textosTextures[i].textura, nil, &sdl.Rect{x, y, largura, altura})

	}

}