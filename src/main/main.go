package main

import (
	"fmt"
	"os"
	"time"

	"ecossistema"
	"tabuleiro"
	"utils"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Projeto para Linguagens de Programacao 2019.02

// AUTOR : LUAN ALVES - 17/0003191
// AUTOR : MARUAN OLIVEIRA - 18/0057685

var (
	window   *sdl.Window
	renderer *sdl.Renderer
	surface  *sdl.Surface
	event    sdl.Event
	font     *ttf.Font
	err      error

	running bool
)

func main() {

	if !Configuracao() {
		os.Exit(1)
	}

	// ESCOPO PRINCIPAL
	utils.Log("logs/logs.txt", "")
	utils.Log("logs/logs.txt", " ------------------ SIMULACAO ------------------ ")

	tb := tabuleiro.TabuleiroNovo("MATRIZ")
	ambienteC := ecossistema.AmbienteNovo()
	ecossistemaC := ecossistema.EcossistemaNovo(ambienteC)

	// Carregar Organismos
	var caminho string = "assets/organismos/"

	ecossistemaC.CarregarOrganismos(caminho)

	ecossistemaC.MapearOrganismos(tb)

	ecossistemaC.ProduzirOxigenio(100000)
	ecossistemaC.ProduzirCarbono(100000)

	running = true
	for running {

		ManipularEventos()

		fmt.Println("---------------- Ciclo :  ", ambienteC.Ciclo(), " --------------------------------")
		time.Sleep(time.Second / 15)
		fmt.Println("")

		tb.Atualizar(surface)

		//if ambienteC.FaseContador() == 0 {

		//tb.Mostrar()

			ecossistemaC.RemoverOrganimosMortos(tb)

			ecossistemaC.LogEcossistema()

		//}

		ecossistemaC.ExecutarCiclo(surface, tb)

		AtualizarTela(ambienteC, ecossistemaC)

		ambienteC.AmbienteFase()

	}

	Encerrar()

	fmt.Println("Fim da Simulação !!!")

}
