package grafico

import "github.com/veandco/go-sdl2/sdl"

func GraficoSequenciador(surface sdl.Surface, alturageral int32, cor uint32, dados Sequenciador) {

	var GX int32 = 510
	var GY int32 = alturageral
	var GLargura int32 = 300
	var GAltura int32 = 100
	var Barra int32 = 2

	var L1 = sdl.Rect{GX, GY, GLargura, Barra}
	surface.FillRect(&L1, 0xFFFF00)

	var L2 = sdl.Rect{GX, GY + GAltura, GLargura, Barra}
	surface.FillRect(&L2, 0xFFFF00)

	var L3 = sdl.Rect{GX, GY, Barra, GAltura}
	surface.FillRect(&L3, 0xFFFF00)

	var L4 = sdl.Rect{GX + GLargura, GY, Barra, GAltura + Barra}
	surface.FillRect(&L4, 0xFFFF00)

	var x1 int32 = GX + Barra + 2

	var ni int = 0
	for ni < 20 {

		var tempc float32 = dados.ValorCorrente()

		if tempc < 0 {
			tempc = 0
		}
		if tempc > 100 {
			tempc = 100
		}

		var L5 = sdl.Rect{x1, GAltura - int32(tempc) + GY, 10, int32(tempc)}
		surface.FillRect(&L5, cor)
		x1 += 15

		ni++
	}

}
