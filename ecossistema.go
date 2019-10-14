package main

import (
	"fmt"
	"strconv"
)

type ecossistema struct {
	contadorprodutor int
	contadoranimal   int
	produtores       (map[string]*produtor)
	consumidores     (map[string]*consumidor)
}

func EcossistemaNovo() *ecossistema {

	p := ecossistema{}
	p.contadoranimal = 0
	p.contadorprodutor = 0

	p.produtores = make(map[string]*produtor)
	p.consumidores = make(map[string]*consumidor)

	return &p
}

func (a *ecossistema) adicionarProdutor(produtor *produtor) {

	a.produtores[strconv.Itoa(a.contadorprodutor)] = produtor

	a.contadorprodutor++
}

func (a *ecossistema) adicionarConsumidor(animalc *consumidor) {

	a.consumidores[strconv.Itoa(a.contadoranimal)] = animalc

	a.contadoranimal++
}

func (a *ecossistema) mapearOrganismos() {
	a.mapearProdutores()
	a.mapearConsumidores()
}

func (a *ecossistema) mapearConsumidores() {

	for p := range a.consumidores {

		var animalc = a.consumidores[p]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		animalc.mudarposicao(x, y)
	}

}

func (a *ecossistema) mapearProdutores() {

	for p := range a.produtores {

		var plantac = a.produtores[p]

		var x int = aleatorionumero(50)
		var y int = aleatorionumero(50)

		plantac.mudarposicao(x, y)
	}

}

func (a *ecossistema) removerOrganimosMortos() {

	for p := range a.produtores {

		var plantac = a.produtores[p]

		if plantac.status() == "morto" {

			fmt.Println("      - Removendo Produtor", p)

			delete(a.produtores, p)
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
