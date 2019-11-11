package main

import (
	"ecossistema"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"grafico"
	"strconv"
)

func AtualizarTela(a *ecossistema.Ambiente, e *Escritor) {

	AtualizarTelaBarraDia(*a, surface)

	//RenderizarTextos(e)

	TextosInfo(a, surface, e)

	grafico.GraficoSequenciador(*surface, 50, 0x782206, a.SeqTemperatura)

	grafico.GraficoSequenciador(*surface, 155, 0x1F4500, a.SeqVento)

	grafico.GraficoSequenciador(*surface, 260, 0x7F4500, a.SeqUmidade)

	grafico.GraficoSequenciador(*surface, 365, 0x5F4500, a.SeqNuvem)

	grafico.GraficoSequenciador(*surface, 470, 0x2F4700, a.SeqLuz)

	grafico.GraficoSequenciador(*surface, 575, 0x3F5F50, a.SeqChuva)

	renderer.Present()

}


func TextosInfo(e *ecossistema.Ambiente, s *sdl.Surface, EscritorC *Escritor) {

	//grafico.RegiaoRetangular(*s, 1, 530, 590,250,0x1F4500)


	EscritorC.EscreveBranco(250, 530,"Informações")

	var alturador int32=550

	EscritorC.EscreveBranco(10, alturador, fmt.Sprintf("Nuvens: %s", e.NuvemCorrente()))
	EscritorC.EscreveBranco(430, alturador, fmt.Sprintf("%d", e.Dia()))

	alturador+=25
	EscritorC.EscreveBranco(10, alturador,  fmt.Sprintf("Temp: %.2f ºC", e.TemperaturaCorrente()))
	EscritorC.EscreveBranco(200, alturador, fmt.Sprintf("Umidade: %s", e.UmidadeCorrente()))
	EscritorC.EscreveBranco(400, alturador, fmt.Sprintf("Chuva : %s", e.ChuvaCorrente()))

	alturador+=25
	EscritorC.EscreveBranco(10, alturador, fmt.Sprintf("Vento: %s", e.VentoCorrente()))
	EscritorC.EscreveBranco(200, alturador, fmt.Sprintf("Luz : %s", e.LuzCorrente()))
	EscritorC.EscreveBranco(400, alturador, fmt.Sprintf("Nuvens: %s", e.NuvemCorrente()))

	alturador+=25
	EscritorC.EscreveBranco(10, alturador, fmt.Sprintf("Vento Origem: %s", e.VentoOrigem()))
	EscritorC.EscreveBranco(200, alturador, fmt.Sprintf("Vento Destino: %s", e.VentoDestino()))
	EscritorC.EscreveBranco(400, alturador, fmt.Sprintf("Vento Modo: %s", e.VentoModo()))


	alturador+=25

	EscritorC.EscreveBranco(10, alturador,  fmt.Sprintf("Chuva Tipo: %s", e.ChuvaTipo()))
	EscritorC.EscreveBranco(200, alturador,  fmt.Sprintf("Chuva Modo: %s", e.ChuvaModo()))
	EscritorC.EscreveBranco(400, alturador,  "Chuva Int.: " + e.ChuvaIntensidadeStatus() )


	alturador+=25

	EscritorC.EscreveBranco(10, alturador,  "Vento Vel.: " + e.VentoVelocidadeCorrenteInfo() )
	EscritorC.EscreveBranco(200, alturador,  "Chuva Int.: " + e.ChuvaIntensidadeInfo() )
	EscritorC.EscreveBranco(400, alturador,  fmt.Sprintf("Ar Seco: %s", e.ArSecoStatus()))

	alturador+=35

	EscritorC.EscreveBranco(10, alturador,  e.SensacaoInfo())

}

func AtualizarTelaBarraDia(a ecossistema.Ambiente, s *sdl.Surface, ) {

	var TamanhoBarra int32=490

	var linhafinal = sdl.Rect{10, 500, TamanhoBarra, 10}
	if a.Fase() == "Dia" {
		s.FillRect(&linhafinal, 0xFFFF00)
	} else {
		s.FillRect(&linhafinal, 0x000080)
	}

	//	var st = a.Sol() * 5
	var t float32=float32(TamanhoBarra)
	var st float32= ((t) / float32(100)) * float32(a.FaseContador())
	var solinha = sdl.Rect{10, 510, int32(st), 10}
	s.FillRect(&solinha, 0xFF4500)

}


func SalvarTela(ambienteC *ecossistema.Ambiente,surface*sdl.Surface,local string){

	var txt string=strconv.Itoa(ambienteC.Ciclo() )
	if len(txt)==1{
		txt="0" + txt
	}
	surface.SaveBMP(local + "current.png")
	surface.SaveBMP(local + "ciclo/" + txt + ".png")


}
