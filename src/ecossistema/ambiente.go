package ecossistema

import (
	"fmt"
	"strconv"

	"utils"

	"github.com/veandco/go-sdl2/sdl"
)

type Ambiente struct {
	fase          string
	faseciclo     int
	fasecontador  int
	dia           int
	sol           int

	ciclo         int
}

func AmbienteNovo() *Ambiente {

	p := Ambiente{}
	p.faseciclo = 100
	p.dia = 0
	p.fase = ""
	p.fasecontador = 0
	p.sol = 0

	p.ciclo = 0

	return &p
}

func (a *Ambiente) ceu() string {

	if a.fase == "Dia" {
		return a.luminosidade(a.sol)
	} else {
		return ""
	}

}

func (a *Ambiente) luminosidade(_sol int) string {

	var _solmodo string = " - "
	if _sol >= 0 && _sol < 20 {
		_solmodo = "Muito Nublado"
	}

	if _sol >= 20 && _sol < 40 {
		_solmodo = "Nublado"
	}

	if _sol >= 40 && _sol < 60 {
		_solmodo = "Normal"
	}

	if _sol >= 60 && _sol < 80 {
		_solmodo = "Ensolarado"
	}

	if _sol >= 80 {
		_solmodo = "Muito Ensolarado"
	}

	return _solmodo

}

func (a *Ambiente) AmbienteFase() {

	// Implementacao FASE - DIA / NOITE

	a.ciclo++

	if a.fase == "" {
		a.fasecontador = a.faseciclo * 2
	}

	if a.fasecontador >= a.faseciclo {
		a.fasecontador = 0
		if a.fase == "Dia" {
			a.fase = "Noite"
			a.sol = 0
			utils.Log("logs.txt", "Noite - "+strconv.Itoa(a.dia)+" [ ]")

		} else {
			a.fase = "Dia"
			a.dia++
			a.sol = utils.Aleatorionumero(100)

			utils.Log("logs.txt", "Dia - "+strconv.Itoa(a.dia)+" [ "+a.ceu()+"]")

		}
	} else {
		a.fasecontador++

		if a.fase == "Dia" {
			modo := utils.Aleatorionumero(100)
			valor := utils.Aleatorionumero(5)

			if modo <= 50 {
				a.sol += valor
			} else {
				a.sol -= valor
			}

		}
	}

	if a.sol < 0 {
		a.sol = a.sol * (-1)
	}

	fmt.Println("")
	fmt.Println("Fase -> ", a.fase)
	fmt.Println("Quantidade de Sol -> ", a.sol)
	fmt.Println("Ceu -> ", a.ceu())

}

func (a *Ambiente) AtualizarTela(s *sdl.Surface,) {

	var linhafinal = sdl.Rect{0, 500, 500, 10}
	if a.Fase() == "Dia" {
		s.FillRect(&linhafinal, 0xFFFF00)
	} else {
		s.FillRect(&linhafinal, 0x000080)
	}

	var st = a.Sol() * 5
	var solinha = sdl.Rect{0, 510, int32(st), 10}
	s.FillRect(&solinha, 0xFF4500)

}

func (a *Ambiente) Fase() string { return a.fase }
func (a *Ambiente) FaseContador() int { return a.fasecontador }
func (a *Ambiente) Ciclo() int { return a.ciclo }
func (a *Ambiente) Sol() int { return a.sol }