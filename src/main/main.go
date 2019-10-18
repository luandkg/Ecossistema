package main

import (
	"fmt"
	"os"
	"time"

	"tabuleiro"
	"ecossistema"
	"utils"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Projeto para Linguagens de Programacao 2019.02

// AUTOR : LUAN ALVES - 17/0003191
// AUTOR : MARUAN OLIVEIRA - 18/0057685

var (
	window          *sdl.Window
	renderer        *sdl.Renderer
	surface         *sdl.Surface
	event           sdl.Event
	font 			*ttf.Font
	err             error

	running         bool
)

func AtualizarTela(a *ecossistema.Ambiente, e *ecossistema.Ecossistema) {

	a.AtualizarTela(surface)

	RenderizarTextos(e)

	renderer.Present()

}

func main() {

	if !Configuracao() {
		os.Exit(1)
	}

	// ESCOPO PRINCIPAL
	utils.Log("logs.txt", "")
	utils.Log("logs.txt", " ------------------ SIMULACAO ------------------ ")

	tb := tabuleiro.TabuleiroNovo("MATRIZ")
	ambienteC := ecossistema.AmbienteNovo()
	ecossistemaC := ecossistema.EcossistemaNovo()

	tb.Limpar()

	// GERAR PRODUTORES
	ecossistemaC.GerarOrganismos("produtor", 10, "Capim Gordura", 200, 100, 300, 0x18B300)
	ecossistemaC.GerarOrganismos("produtor", 10, "Capim Verde", 300, 150, 600, 0x89FF77)
	ecossistemaC.GerarOrganismos("produtor", 10, "Laranjeira", 500, 200, 10000, 0xC2FFDC)
	ecossistemaC.GerarOrganismos("produtor", 10, "Ervacidreira", 300, 300, 1000, 0x0EB355)

	// GERAR CONSUMIDORES
	ecossistemaC.GerarOrganismos("consumidor", 10, "Rato", 300, 200, 3000, 0xCC2700)
	ecossistemaC.GerarOrganismos("consumidor", 4, "Roedor", 400, 200, 5000, 0xFF845F)
	ecossistemaC.GerarOrganismos("consumidor", 6, "Coelho", 500, 250, 8000, 0xFF4570)

	ecossistemaC.MapearOrganismos(tb)

	running = true
	for running {

		ManipularEventos()

		fmt.Println("---------------- Ciclo :  ", ambienteC.Ciclo(), " --------------------------------")
		time.Sleep(time.Second / 4)
		fmt.Println("")

		tb.Atualizar(surface)

		if ambienteC.FaseContador() == 0 {

			ecossistemaC.RemoverOrganimosMortos()

			ecossistemaC.LogEcossistema()

		}

		ecossistemaC.ExecutarCiclo(surface, tb)

		AtualizarTela(ambienteC, ecossistemaC)

		ambienteC.AmbienteFase()

	}

	Encerrar()

	fmt.Println("Fim da Simulação !!!")

}
