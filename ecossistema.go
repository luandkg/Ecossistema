package main

import (
	"fmt"
	"strconv"
)

type ecossistema struct {
	contadorplanta int
	contadoranimal int
	plantas        (map[string]*planta)
	consumidores   (map[string]*consumidor)
}

func EcossistemaNovo() *ecossistema {

	p := ecossistema{}
	p.contadoranimal = 0
	p.contadorplanta = 0

	p.plantas = make(map[string]*planta)
	p.consumidores = make(map[string]*consumidor)

	return &p
}

func (a *ecossistema) adicionarPlanta(plantac *planta) {

	a.plantas[strconv.Itoa(a.contadorplanta)] = plantac

	a.contadorplanta++
}

func (a *ecossistema) adicionarConsumidor(animalc *consumidor) {

	a.consumidores[strconv.Itoa(a.contadoranimal)] = animalc

	a.contadoranimal++
}

func (a *ecossistema) mapearConsumidores() {

	for p := range a.consumidores {

		var animalc = a.consumidores[p]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		animalc.mudarposicao(x, y)
	}

}

func (a *ecossistema) mapearPlantas() {

	for p := range a.plantas {

		var plantac = a.plantas[p]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		plantac.mudarposicao(x, y)
	}

}

func (a *ecossistema) removerOrganimosMortos() {

	for p := range a.plantas {

		var plantac = a.plantas[p]

		if plantac.status() == "morto" {

			fmt.Println("      - Removendo Planta", p)

			delete(a.plantas, p)
		}

	}

	for p := range a.consumidores {

		var animalc = a.consumidores[p]

		if animalc.status() == "morto" {

			fmt.Println("      - Removendo consumidor", p)

			delete(a.consumidores, p)
		}

	}

}
