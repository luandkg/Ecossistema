package ecossistema

import (
	"fmt"
	"strconv"
	"tabuleiro"

	"bioxml"
	"github.com/veandco/go-sdl2/sdl"
	"utils"
)

type Ecossistema struct {
	contadorProdutor   int
	contadorConsumidor int
	produtores         (map[string]*Produtor)
	consumidores       (map[string]*Consumidor)

	AmbienteC *Ambiente

	produtores_contagem   int
	consumidores_contagem int
	produtores_mortos     int
	consumidores_mortos   int
}

func EcossistemaNovo(ambienteC *Ambiente) *Ecossistema {

	p := Ecossistema{}
	p.contadorConsumidor = 0
	p.contadorProdutor = 0

	p.produtores = make(map[string]*Produtor)
	p.consumidores = make(map[string]*Consumidor)

	p.AmbienteC = ambienteC

	return &p
}

func (e *Ecossistema) AdicionarProdutor(produtorC *Produtor) {

	produtorC._nomecompleto = produtorC._nome + "::" + strconv.Itoa(e.contadorProdutor)

	e.produtores[strconv.Itoa(e.contadorProdutor)] = produtorC

	fmt.Println("       --- Produtor : ", produtorC.NomeCompleto(), " Nasceu !!!")

	e.contadorProdutor++
	e.produtores_contagem++

}

func (e *Ecossistema) AdicionarConsumidor(consumidorC *Consumidor) {

	consumidorC._nomecompleto = consumidorC._nome + "::" + strconv.Itoa(e.contadorConsumidor)

	e.consumidores[strconv.Itoa(e.contadorConsumidor)] = consumidorC

	fmt.Println("       --- Consumidor : ", consumidorC.NomeCompleto(), " Nasceu !!!")

	e.contadorConsumidor++
	e.consumidores_contagem++

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

	for p := range e.produtores {

		var produtorCorrente = e.produtores[p]

		if produtorCorrente.Status() == "morto" {

			fmt.Println("      - Removendo Produtor : ", produtorCorrente.NomeCompleto())

			peca := tb.RecuperarPeca(produtorCorrente.x(), produtorCorrente.y())

			peca.LiberarPosicao()

			delete(e.produtores, p)

			e.produtores_mortos++

		}

	}

	for p := range e.consumidores {

		var consumidorCorrente = e.consumidores[p]

		if consumidorCorrente.Status() == "morto" {

			fmt.Println("      - Removendo consumidor : ", consumidorCorrente.NomeCompleto())

			peca := tb.RecuperarPeca(consumidorCorrente.x(), consumidorCorrente.y())

			peca.LiberarPosicao()

			delete(e.consumidores, p)

			e.consumidores_mortos++

		}

	}

}

func (e *Ecossistema) ExecutarCiclo(surface *sdl.Surface, tb *tabuleiro.Tabuleiro) {

	e.executarCicloProdutores(surface)

	e.executarCicloConsumidores(surface, tb)

}

func (e *Ecossistema) executarCicloProdutores(surface *sdl.Surface) {

	fmt.Println("PRODUTORES")

	for p := range e.produtores {

		produtorc := e.produtores[p]

		if produtorc.Status() == "vivo" {

			//	fmt.Println("      - ", produtorc.toString())
			produtorc.vivendo()
			produtorc.atualizar(surface)

		}

	}

}

func (e *Ecossistema) executarCicloConsumidores(surface *sdl.Surface, tb *tabuleiro.Tabuleiro) {

	fmt.Println("CONSUMIDORES")

	for p := range e.consumidores {

		consumidorc := e.consumidores[p]

		if consumidorc.Status() == "vivo" {

			//fmt.Println("      - ", consumidorc.toString())
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

	utils.Log("logs/logs.txt", "-------------------------------------------------------------------------")

	utils.Log("logs/logs.txt", "Plantas - "+strconv.Itoa(len(e.produtores)))

	utils.Log("logs/logs.txt", "Consumidores - "+strconv.Itoa(len(e.consumidores)))

	utils.Log("logs/logs.txt", "Gas Oxigenio - "+fmt.Sprintf("%f", e.AmbienteC.gasOxigenio))
	utils.Log("logs/logs.txt", "Gas Carbonico - "+fmt.Sprintf("%f", e.AmbienteC.gasCarbonico))

	utils.Log("logs/logs.txt", "Energia - "+fmt.Sprintf("%f", e.EnergiaTotal()))

	utils.Log("logs/logs.txt", "CONTAGEM")
	utils.Log("logs/logs.txt", "    - Produtores Contagem - "+strconv.Itoa((e.produtores_contagem)))
	utils.Log("logs/logs.txt", "    - Produtores Vivos - "+strconv.Itoa((e.Produtores_Vivos())))
	utils.Log("logs/logs.txt", "    - Produtores Mortos - "+strconv.Itoa((e.produtores_mortos)))

	utils.Log("logs/logs.txt", "    - Consumidores Contagem - "+strconv.Itoa((e.consumidores_contagem)))
	utils.Log("logs/logs.txt", "    - Consumidores Vivos - "+strconv.Itoa((e.Consumidores_Vivos())))
	utils.Log("logs/logs.txt", "    - Consumidores Mortos - "+strconv.Itoa((e.consumidores_mortos)))

}

func (e *Ecossistema) EnergiaTotal() float32 {

	var energia float32 = 0

	for p := range e.produtores {

		var pcorrente = e.produtores[p]

		if pcorrente.Status() == "vivo" {

			energia += pcorrente._energia
		}

	}

	return energia
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

func (e *Ecossistema) GerarOrganismos(tipo string, quantidade int, nome string, adulto int, reproducao int, gestacao int, vida int, cor uint32) {

	switch tipo {

	case "produtor":
		for i := 0; i < quantidade; i++ {
			e.AdicionarProdutor(ProdutorNovo(nome, adulto, reproducao, gestacao, vida, cor, e))
		}

	case "consumidor":
		for i := 0; i < quantidade; i++ {
			e.AdicionarConsumidor(ConsumidorNovo(nome, adulto, reproducao, vida, cor, e))
		}

	}

}

func (e *Ecossistema) ProduzirOxigenio(valor float64) {
	e.AmbienteC.gasOxigenio += valor
}

func (e *Ecossistema) ProduzirCarbono(valor float64) {
	e.AmbienteC.gasCarbonico += valor
}

func (e *Ecossistema) CarregarOrganismos(caminho string) {

	for _, OrganismoNome := range bioxml.ListarOrganismos(caminho) {

		var organismoC *bioxml.Organismo = bioxml.CarregarOrganismo(caminho + OrganismoNome + ".organismo")

		if organismoC.Base.Tipo == "Produtor" {
			var cor uint32 = organismoC.Base.Cor

			e.GerarOrganismos("produtor", 10, OrganismoNome, organismoC.Base.Adulto, organismoC.Reproducao.Frequencia, organismoC.Reproducao.Gestacao, organismoC.Base.Vida, cor)
		}
		if organismoC.Base.Tipo == "Consumidor" {
			var cor uint32 = organismoC.Base.Cor

			e.GerarOrganismos("consumidor", 10, OrganismoNome, organismoC.Base.Adulto, organismoC.Reproducao.Frequencia, organismoC.Reproducao.Gestacao, organismoC.Base.Vida, cor)
		}
	}

}

func (e *Ecossistema) Produtores_Contagem() int { return e.produtores_contagem }
func (e *Ecossistema) Produtores_Mortos() int   { return e.produtores_mortos }
func (e *Ecossistema) Produtores_Vivos() int    { return e.produtores_contagem - e.produtores_mortos }

func (e *Ecossistema) Consumidores_Contagem() int { return e.consumidores_contagem }
func (e *Ecossistema) Consumidores_Mortos() int   { return e.consumidores_mortos }
func (e *Ecossistema) Consumidores_Vivos() int    { return e.consumidores_contagem - e.consumidores_mortos }
