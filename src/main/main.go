package main

import (
	"fmt"
	"os"
	"time"

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

func AtualizarTela(ecossistemaC *ecossistema.Ecossistema) {

	RenderizarTextos(ecossistemaC)

	renderer.Present()

}

func main() {

	if !Configuracao() {
		os.Exit(1)
	}

	// ESCOPO PRINCIPAL
	utils.Log("logs.txt", "")
	utils.Log("logs.txt", " ------------------ SIMULACAO ------------------ ")

	tb := Tabuleiro_novo("MATRIZ")
	ambienteC := ecossistema.AmbienteNovo()
	ecossistemaC := ecossistema.EcossistemaNovo()

	tb.limpar()

	// PLANTAS

	for i := 0; i < 10; i++ {
		ecossistemaC.AdicionarProdutor(ecossistema.Plantanovo("Capim Gordura", 200, 100, 300, 0xADFF2F, ecossistemaC))
	}
	for i := 0; i < 10; i++ {
		ecossistemaC.AdicionarProdutor(ecossistema.Plantanovo("Capim Verde", 300, 150, 600, 0x808000, ecossistemaC))
	}
	for i := 0; i < 10; i++ {
		ecossistemaC.AdicionarProdutor(ecossistema.Plantanovo("Laranjeira", 500, 200, 10000, 0xDAA520, ecossistemaC))
	}
	for i := 0; i < 10; i++ {
		ecossistemaC.AdicionarProdutor(ecossistema.Plantanovo("Ervacidreira", 300, 300, 1000, 0xFFFF00, ecossistemaC))
	}

	// ANIMAIS

	for i := 0; i < 10; i++ {
		ecossistemaC.AdicionarConsumidor(ecossistema.ConsumidorNovo("Rato", 200, 200, 2000, 0xDDA0DD, ecossistemaC))
	}

	for i := 0; i < 4; i++ {
		ecossistemaC.AdicionarConsumidor(ecossistema.ConsumidorNovo("Roeador", 400, 200, 5000, 0xEE82EE, ecossistemaC))
	}

	for i := 0; i < 6; i++ {
		ecossistemaC.AdicionarConsumidor(ecossistema.ConsumidorNovo("Coelho", 500, 250, 8000, 0x7B68EE, ecossistemaC))
	}

	ecossistemaC.MapearOrganismos()

	running = true
	for running {

		ManipularEventos()

		fmt.Println("---------------- Ciclo :  ", ambienteC.Ciclo(), " --------------------------------")
		time.Sleep(time.Second / 4)
		fmt.Println("")

		tb.atualizar(surface, ambienteC)

		if ambienteC.FaseContador() == 0 {

			ecossistemaC.RemoverOrganimosMortos()

			ecossistemaC.LogEcossistema()

		}

		ecossistemaC.ExecutarCiclo(surface)

		AtualizarTela(ecossistemaC)

		ambienteC.AmbienteFase()

	}

	Encerrar()

	fmt.Println("Fim da Simulação !!!")

}
