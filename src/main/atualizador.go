package main

import (
	"ecossistema"
	"github.com/veandco/go-sdl2/sdl"
	"grafico"
)

func AtualizarTela(a *ecossistema.Ambiente, e *ecossistema.Ecossistema) {

	AtualizarTelaBarraDia(*a, surface)

	RenderizarTextos(e)

	grafico.GraficoSequenciador(*surface, 50, 0x782206, e.AmbienteC.SeqTemperatura)

	grafico.GraficoSequenciador(*surface, 155, 0x1F4500, e.AmbienteC.SeqVento)

	grafico.GraficoSequenciador(*surface, 260, 0x7F4500, e.AmbienteC.SeqUmidade)

	grafico.GraficoSequenciador(*surface, 365, 0x5F4500, e.AmbienteC.SeqNuvem)

	grafico.GraficoSequenciador(*surface, 470, 0x2F4700, e.AmbienteC.SeqLuz)

	grafico.GraficoSequenciador(*surface, 575, 0x3F5F50, e.AmbienteC.SeqChuva)

	renderer.Present()

}

func AtualizarTelaBarraDia(a ecossistema.Ambiente, s *sdl.Surface, ) {

	var linhafinal = sdl.Rect{0, 500, 500, 10}
	if a.Fase() == "Dia" {
		s.FillRect(&linhafinal, 0xFFFF00)
	} else {
		s.FillRect(&linhafinal, 0x000080)
	}

	//	var st = a.Sol() * 5
	var st = (510 / 100) * a.FaseContador()
	var solinha = sdl.Rect{0, 510, int32(st), 10}
	s.FillRect(&solinha, 0xFF4500)

}
