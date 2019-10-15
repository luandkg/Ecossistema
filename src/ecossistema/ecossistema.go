package ecossistema

import (
	"fmt"
	"strconv"
	"utils"
)

type Ecossistema struct {
	contadorprodutor int
	contadoranimal   int
	produtores       (map[string]*produtor)
	consumidores     (map[string]*consumidor)
}

func EcossistemaNovo() *Ecossistema {

	p := Ecossistema{}
	p.contadoranimal = 0
	p.contadorprodutor = 0

	p.produtores = make(map[string]*produtor)
	p.consumidores = make(map[string]*consumidor)

	return &p
}

func (a *Ecossistema) AdicionarProdutor(produtor *produtor) {

	a.produtores[strconv.Itoa(a.contadorprodutor)] = produtor

	a.contadorprodutor++
}

func (a *Ecossistema) AdicionarConsumidor(animalc *consumidor) {

	a.consumidores[strconv.Itoa(a.contadoranimal)] = animalc

	a.contadoranimal++
}

func (a *Ecossistema) MapearOrganismos() {
	a.MapearProdutores()
	a.MapearConsumidores()
}

func (a *Ecossistema) MapearConsumidores() {

	for p := range a.consumidores {

		var animalc = a.consumidores[p]

		var x int = utils.Aleatorionumero(50)
		var y int = utils.Aleatorionumero(50)

		animalc.mudarposicao(x, y)
	}

}

func (a *Ecossistema) MapearProdutores() {

	for p := range a.produtores {

		var plantac = a.produtores[p]

		var x int = utils.Aleatorionumero(50)
		var y int = utils.Aleatorionumero(50)

		plantac.mudarposicao(x, y)
	}

}

func (a *Ecossistema) RemoverOrganimosMortos() {

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
