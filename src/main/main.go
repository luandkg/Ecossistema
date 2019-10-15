package main

import (
	"fmt"
	"os"
	"strconv"
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
		ecossistemaC.AdicionarConsumidor(ecossistema.Consumidor("Rato", 200, 200, 2000, 0xDDA0DD, ecossistemaC))
	}

	for i := 0; i < 4; i++ {
		ecossistemaC.AdicionarConsumidor(ecossistema.Consumidor("Roeador", 400, 200, 5000, 0xEE82EE, ecossistemaC))
	}

	for i := 0; i < 6; i++ {
		ecossistemaC.AdicionarConsumidor(ecossistema.Consumidor("Coelho", 500, 250, 8000, 0x7B68EE, ecossistemaC))
	}

	ecossistemaC.MapearOrganismos()

	running = true
	for running {

		ManipularEventos()

		fmt.Println("---------------- Ciclo :  ", ambienteC.GetCiclo(), " --------------------------------")
		time.Sleep(time.Second / 4)
		fmt.Println("")

		fmt.Println("PRODUTORES")

		tb.atualizar(surface, ambienteC)

		for p := range ecossistemaC.produtores {
			produtorc := ecossistemaC.produtores[p]

			if produtorc.Status() == "vivo" {

				produtorc._nomecompleto = produtorc._nome + " " + p
				fmt.Println("      - ", produtorc.toString())
				produtorc.vivendo()
				produtorc.atualizar(surface)

			}

		}

		fmt.Println("CONSUMIDORES")

		for p := range ecossistemaC.consumidores {

			consumidorc := ecossistemaC.consumidores[p]

			if consumidorc.status() == "vivo" {

				fmt.Println("      - ", consumidorc.toString())
				consumidorc.vivendo()
				consumidorc.movimento()
				consumidorc.atualizar(surface)

			}

		}

		AtualizarTela(ecossistemaC)

		fmt.Println("")

		ambienteC.ambiente()

		fmt.Println("Fase -> ", ambienteC.fase)
		fmt.Println("Quantidade de Sol -> ", ambienteC.sol)
		fmt.Println("Ceu -> ", ambienteC.ceu())

		if ambienteC.fasecontador == 0 {

			utils.Log("logs.txt", "Plantas - "+strconv.Itoa(len(ecossistemaC.produtores)))
			utils.Log("logs.txt", "Consumidores - "+strconv.Itoa(len(ecossistemaC.consumidores)))

			ecossistemaC.RemoverOrganimosMortos()
		}
	}

	Encerrar()

	fmt.Println("Fim da Simulação !!!")

}
