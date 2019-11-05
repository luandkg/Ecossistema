package ecossistema

import (
	"fmt"
	"strconv"
	"tabuleiro"

	"utils"

	"github.com/veandco/go-sdl2/sdl"
)

type Ecossistema struct {
	contadorprodutor int
	contadoranimal   int
	produtores       (map[string]*Produtor)
	consumidores     (map[string]*Consumidor)

	ambienteC *Ambiente
}

func EcossistemaNovo(ambienteC*Ambiente) *Ecossistema {

	p := Ecossistema{}
	p.contadoranimal = 0
	p.contadorprodutor = 0

	p.produtores = make(map[string]*Produtor)
	p.consumidores = make(map[string]*Consumidor)

	p.ambienteC = ambienteC

	return &p
}

func (e *Ecossistema) AdicionarProdutor(produtor *Produtor) {

	e.produtores[strconv.Itoa(e.contadorprodutor)] = produtor

	e.contadorprodutor++
}

func (e *Ecossistema) AdicionarConsumidor(animalc *Consumidor) {

	e.consumidores[strconv.Itoa(e.contadoranimal)] = animalc

	e.contadoranimal++
}

func (e *Ecossistema) MapearOrganismos(tb *tabuleiro.Tabuleiro) {
	e.MapearProdutores(tb)
	e.MapearConsumidores(tb)
}

func (e *Ecossistema) MapearConsumidores(tb *tabuleiro.Tabuleiro) {

	for p := range e.consumidores {

		var animalc = e.consumidores[p]

		var posicaoOcupada = true

		var x = 0
		var y = 0

		for posicaoOcupada {

			x = utils.Aleatorionumero(50)
			y = utils.Aleatorionumero(50)

			peca := tb.RecuperarPeca(x, y)

			posicaoOcupada = peca.VerificarPosicao()

			if posicaoOcupada == false {
				animalc.mudarposicao(x, y)
				peca.OcuparPosicao()
			}

		}

	}

}

func (e *Ecossistema) MapearProdutores(tb *tabuleiro.Tabuleiro) {

	for p := range e.produtores {

		var plantac = e.produtores[p]

		var posicaoOcupada = true

		var x = 0
		var y = 0

		for posicaoOcupada {

			x = utils.Aleatorionumero(50)
			y = utils.Aleatorionumero(50)

			peca := tb.RecuperarPeca(x, y)

			posicaoOcupada = peca.VerificarPosicao()

			if posicaoOcupada == false {
				plantac.mudarposicao(x, y)
				peca.OcuparPosicao()
			}

		}

	}

}

func (e *Ecossistema) RemoverOrganimosMortos(tb *tabuleiro.Tabuleiro) {

	for index, plantac := range e.produtores {

		if plantac.Status() == "morto" {

			fmt.Println("      - Removendo Produtor", index)

			peca := tb.RecuperarPeca(plantac.x(), plantac.y())

			peca.LiberarPosicao()

			delete(e.produtores, index)
		}

	}

	for p := range e.consumidores {

		var animalc = e.consumidores[p]

		if animalc.Status() == "morto" {

			fmt.Println("      - Removendo consumidor", p)

			delete(e.consumidores, p)
		}

	}

}

func (e *Ecossistema) ExecutarCiclo (surface *sdl.Surface, tb *tabuleiro.Tabuleiro) {

	e.executarCicloProdutores(surface)

	e.executarCicloConsumidores(surface, tb)

}

func (e *Ecossistema) executarCicloProdutores (surface *sdl.Surface) {

	fmt.Println("PRODUTORES")

	for p := range e.produtores {

		produtorc := e.produtores[p]

		if produtorc.Status() == "vivo" {

			produtorc._nomecompleto = produtorc._nome + " " + p
			fmt.Println("      - ", produtorc.toString())
			produtorc.vivendo()
			produtorc.atualizar(surface)

		}

	}

}

func (e *Ecossistema) executarCicloConsumidores (surface *sdl.Surface, tb *tabuleiro.Tabuleiro) {

	fmt.Println("CONSUMIDORES")

	for p := range e.consumidores {

		consumidorc := e.consumidores[p]

		if consumidorc.Status() == "vivo" {

			fmt.Println("      - ", consumidorc.toString())
			consumidorc.vivendo(tb)

			if consumidorc.TemAlvo() {

				consumidorc.CacarAlvo(tb)

			} else {

				consumidorc.Movimento(tb)

			}

			consumidorc.VerificarAlvo(tb)

			consumidorc.atualizar(surface)

		}

	}

}

func (e *Ecossistema) LogEcossistema() {

	utils.Log("logs.txt", "Plantas - "+strconv.Itoa(len(e.produtores)))

	utils.Log("logs.txt", "Consumidores - "+strconv.Itoa(len(e.consumidores)))

	s1 :=fmt.Sprintf("%f",e.ambienteC.gasOxigenio)
	s2:=fmt.Sprintf("%f",e.ambienteC.gasCarbonico)

	utils.Log("logs.txt", "Gas Oxigenio - "+ s1)
	utils.Log("logs.txt", "Gas Carbonico - "+ s2)

}

func (e *Ecossistema) TotalProdutores() int {

	return len(e.produtores)

}

func (e *Ecossistema) TotalProdutoresFase() (int, int) {

	var contadorJovem = 0
	var contadorAdulto = 0

	for _, produtor := range e.produtores {

		switch produtor.Fase() {

		case "nascido":
			contadorJovem += 1

		case "adulto":
			contadorAdulto += 1

		}

	}

	return contadorJovem, contadorAdulto

}

func (e *Ecossistema) TotalConsumidores() int {

	return len(e.consumidores)

}

func (e *Ecossistema) TotalConsumidoresFase() (int, int) {

	var contadorJovem = 0
	var contadorAdulto = 0

	for _, consumidor := range e.consumidores {

		switch consumidor.Fase() {

		case "nascido":
			contadorJovem += 1

		case "adulto":
			contadorAdulto += 1

		}

	}

	return contadorJovem, contadorAdulto

}

func (e *Ecossistema) GerarOrganismos (tipo string, quantidade int, nome string, adulto int, reproducao int, vida int, cor uint32, vCadeiaAlimentar int) {

	switch tipo {

	case "produtor":
		for i := 0; i < quantidade; i++ {
			e.AdicionarProdutor(ProdutorNovo(nome, adulto, reproducao, vida, cor, e))
		}

	case "consumidor":
		for i := 0; i < quantidade; i++ {
			e.AdicionarConsumidor(ConsumidorNovo(nome, adulto, reproducao, vida, cor, e, vCadeiaAlimentar))
		}

	}

}

func (e*Ecossistema) produzirOxigenio (valor float32) {

	e.ambienteC.gasOxigenio+=valor

}

func (e*Ecossistema) produzirCarbono (valor float32) {

	e.ambienteC.gasCarbonico+=valor

}