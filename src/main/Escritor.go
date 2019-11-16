package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Escritor struct{

	renderizador *sdl.Renderer
	font *ttf.Font
	_fontetamanho int

}

type textobloco struct{
	largura int32
	altura int32
	textura *sdl.Texture
}


func EscritorNovo(fontecaminho string,tamanho int,renderizador *sdl.Renderer) *Escritor {

	ret := Escritor{}
	ret.renderizador=renderizador
	ret._fontetamanho=tamanho

	if ret.font, err = ttf.OpenFont(fontecaminho, tamanho); err != nil {
		fmt.Printf("Failed to open font: %s\n", err)
	}

	fmt.Println("Fonte Carregada : ",ret.font.FaceFamilyName())


	return &ret
}
func(e*Escritor) EscreveCor(posx int32,posy int32,t string,cor sdl.Color) {


	var solidSurface *sdl.Surface
	if solidSurface, err = e.font.RenderUTF8Solid(t, cor); err != nil {
		fmt.Printf("Failed to render text: %s\n", err)
	}

	var novoTextoTextura *sdl.Texture
	var textoLargura=0
	var textoAltura=0

	if novoTextoTextura, err = e.renderizador.CreateTextureFromSurface(solidSurface); err != nil {
		fmt.Printf("Failed to create texture: %s\n", err)
	}

	if textoLargura, textoAltura, err = e.font.SizeUTF8(t); err != nil {
		fmt.Printf("Failed to get width or height: %s\n", err)
	}



	var textoRenderizado = textobloco{int32(textoLargura), int32(textoAltura), novoTextoTextura}


	e.renderizador.Copy(textoRenderizado.textura, nil, &sdl.Rect{posx, posy, textoRenderizado.largura, textoRenderizado.altura})



	solidSurface.Free()

}

func(e*Escritor) EscreveBranco(posx int32,posy int32,t string) {

	e.EscreveCor(posx,posy,t,sdl.Color{255, 255, 255, 255})

}

func(e*Escritor) EscrevePreto(posx int32,posy int32,t string) {

	e.EscreveCor(posx,posy,t,sdl.Color{0, 0, 0, 0})

}


func(p*Escritor) FonteNome()string{return p.font.FaceFamilyName() }
func(p*Escritor) FonteTamanho()int{return p._fontetamanho }
