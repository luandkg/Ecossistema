package main

import "C"
import (
	"ecossistema"
	"fmt"
	"os"
	"tabuleiro"
	"time"

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

	dir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Diretorio nao encontrado: %s\n", err)
	}

	var LOCAL_FONTE string = dir + "/assets/fonts/OpenSans-Regular.ttf"
	var LOCAL_ORGANISMOS string = "assets/organismos/"
	var LOCAL_PRINTS string = "assets/prints/"

	EscritorC := EscritorNovo(LOCAL_FONTE, 14, renderer)

	tb := tabuleiro.TabuleiroNovo("MATRIZ")
	ambienteC := ecossistema.AmbienteNovo()
	ecossistemaC := ecossistema.EcossistemaNovo(ambienteC)

	ecossistemaC.CarregarOrganismos(LOCAL_ORGANISMOS)

	ecossistemaC.MapearOrganismos(tb)

	ecossistemaC.ProduzirOxigenio(100000)
	ecossistemaC.ProduzirCarbono(100000)

	running = true
	for running {

		ManipularEventos()

		fmt.Println("---------------- Ciclo :  ", ambienteC.Ciclo(), " --------------------------------")
		time.Sleep(time.Second / 5)
		fmt.Println("")

		tb.Atualizar(surface)

		ecossistemaC.RemoverOrganimosMortos(tb)

		ecossistemaC.LogEcossistema()

		ecossistemaC.ExecutarCiclo(surface, tb)

		AtualizarTela(ambienteC, EscritorC)

		ambienteC.AmbienteFase()

		SalvarTela(ambienteC, surface, LOCAL_PRINTS)

	}

	Encerrar()

	fmt.Println("Fim da Simulação !!!")

}
