package main

import (
	"strconv"
)

type ambiente struct {
	fase         string
	faseciclo    int
	fasecontador int
	dia          int
	sol          int

	ciclo int
}

func AmbienteNovo() *ambiente {

	p := ambiente{}
	p.faseciclo = 100
	p.dia = 0
	p.fase = ""
	p.fasecontador = 0
	p.sol = 0

	p.ciclo = 0

	return &p
}

func (a *ambiente) ceu() string {

	if a.fase == "Dia" {
		return a.luminosidade(a.sol)
	} else {
		return ""
	}

}

func (a *ambiente) luminosidade(_sol int) string {

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

func (a *ambiente) ambiente() {

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
			log("logs.txt", "Noite - "+strconv.Itoa(a.dia)+" [ ]")

		} else {
			a.fase = "Dia"
			a.dia++
			a.sol = aleatorionumero(100)

			log("logs.txt", "Dia - "+strconv.Itoa(a.dia)+" [ "+a.ceu()+"]")

		}
	} else {
		a.fasecontador++

		if a.fase == "Dia" {
			modo := aleatorionumero(100)
			valor := aleatorionumero(5)

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
}
