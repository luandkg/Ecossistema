package main

import (
	"math/rand"
	"strconv"
	"time"
)

// Variaveis contagem de CICLO
var fase string = ""
var faseciclo int = 10
var fasecontador int = 0
var dia int = 0

// Variaveis Qualificadores do Ambiente
var sol int = 0
var solmodo string = " - "

func luminosidade(_sol int) string {

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
		_solmodo = "Ensolado"
	}

	if _sol >= 80 {
		_solmodo = "Muito Ensolado"
	}

	return _solmodo

}

func ambiente() {

	// Implementacao FASE - DIA / NOITE

	if fase == "" {
		fasecontador = faseciclo * 2
	}

	if fasecontador >= faseciclo {
		fasecontador = 0
		if fase == "Dia" {
			fase = "Noite"
			sol = 0
			solmodo = " - "
			log("logs.txt", "Noite - "+strconv.Itoa(dia)+" [ ]")

		} else {
			fase = "Dia"
			dia++
			r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
			sol = r1.Intn(100)

			solmodo = luminosidade(sol)

			log("logs.txt", "Dia - "+strconv.Itoa(dia)+" [ "+solmodo+"]")

		}
	} else {
		fasecontador++
	}

}
