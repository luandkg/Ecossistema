package ecossistema

import (
	"bioxml"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"tabuleiro"
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

	doneProdutores := make(chan bool)

	doneConsumidores := make(chan bool)

	go e.executarCicloProdutores(surface, doneProdutores)

	go e.executarCicloConsumidores(surface, tb, doneConsumidores)

	<-doneProdutores

	<-doneConsumidores

	close(doneProdutores)

	close(doneConsumidores)

}

func (e *Ecossistema) executarCicloProdutores(surface *sdl.Surface, done chan bool) {

	fmt.Println("PRODUTORES")

	var wg sync.WaitGroup

	max := 10

	sem := make(chan bool, max)

	for p := range e.produtores {

		wg.Add(1)

		sem <- true

		produtorc := e.produtores[p]

		go func() {

			if produtorc.Status() == "vivo" {

				e.FatoresAbioticosDeReproducaoProdutores(produtorc)
				e.FatoresDeSobrevivenciaProdutores(produtorc)
				produtorc.atualizar(surface)

				//	fmt.Println("      - ", produtorc.toString())
				produtorc.vivendo()

			}

			wg.Done()

			<-sem

		}()

	}

	fmt.Println("goroutiness ----------------------- : ", runtime.NumGoroutine())


	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	close(sem)

	wg.Wait()

	done <- true

}

func (e *Ecossistema) executarCicloConsumidores(surface *sdl.Surface, tb *tabuleiro.Tabuleiro, done chan bool) {

	fmt.Println("CONSUMIDORES")

	var wg sync.WaitGroup

	max := 10

	sem := make(chan bool, max)

	for p := range e.consumidores {

		wg.Add(1)

		sem <- true

		consumidorc := e.consumidores[p]

		go func() {

			if consumidorc.Status() == "vivo" {

				//fmt.Println("      - ", consumidorc.toString())
				consumidorc.vivendo(tb)

				e.FatoresAbioticosDeLuz(tb, consumidorc)
				e.FatoresDeSobrevivenciaConsumidores(consumidorc)

				consumidorc.atualizar(surface)
			}

			wg.Done()

			<-sem

		}()

	}

	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	close(sem)

	wg.Wait()

	done <- true

}


func (e *Ecossistema) FatoresAbioticosDeLuz (tb *tabuleiro.Tabuleiro, consumidorc *Consumidor) {

	//e.AmbienteC.LuzCorrenteValor() > 20

	if consumidorc._comportamento == "Noturno" {

		if e.AmbienteC.LuzCorrenteValor() < 20 { //Verifica a claridade do dia ; SE < 20 = Escuro;

			if consumidorc.TemAlvo() {
				consumidorc.CacarAlvo(tb)

			} else {
				consumidorc.Movimento(tb)

			}
			consumidorc.VerificarAlvo(tb)
		}

	}

	if consumidorc._comportamento == "Diurno" {

		if e.AmbienteC.LuzCorrenteValor() > 20 { //Verifica a claridade do dia ; SE < 20 = Escuro;

			if consumidorc.TemAlvo() {
			consumidorc.CacarAlvo(tb)

			} else {
			consumidorc.Movimento(tb)

			}
			consumidorc.VerificarAlvo(tb)
		}
	}

}

func (e *Ecossistema) FatoresDeSobrevivenciaConsumidores ( consumidorc *Consumidor) {

	if e.AmbienteC.TemperaturaCorrente() < consumidorc._temperaturaMin || e.AmbienteC.TemperaturaCorrente() > consumidorc._temperaturaMax {

		if randomBool() == true {
			fmt.Println("===========>>>CONSUMIDOR MORREU POR FATORES DE TEMPERATURA: "+fmt.Sprintf("%f", e.AmbienteC.TemperaturaCorrente()))
			consumidorc.morrer()
		}

	} else {

		if e.AmbienteC.VentoCorrenteValor() > consumidorc._ventoMax || e.AmbienteC.ChuvaModo() == consumidorc._morrePorChuvaEspecial {

			if randomBool() == true {
				fmt.Println("=========>>> CONSUMIDOR MORREU POR FATORES DE VENTO OU CHUVA ESPECIAL: "+fmt.Sprintf("%f", e.AmbienteC.VentoCorrenteValor()))
				consumidorc.morrer()
			}

		}

	}

}

func (e *Ecossistema) FatoresDeSobrevivenciaProdutores (produtorc *Produtor) {

	if e.AmbienteC.TemperaturaMedia() < produtorc._temperaturaMin || e.AmbienteC.TemperaturaMedia() > produtorc._temperaturaMax {

		if randomBool() {
			fmt.Println("============================ PRODUTOR MORREU POR FATORES DE TEMPERATURA: "+fmt.Sprintf("%f", e.AmbienteC.TemperaturaMedia()))
			produtorc.morrer()

		}

	} else {

		if e.AmbienteC.UmidadeCorrenteValor() < produtorc._umidadeMin || e.AmbienteC.UmidadeCorrenteValor() > produtorc._umidadeMax {

			if randomBool() {
				fmt.Println("============================PRODUTOR MORREU POR FATORES DE {UMIDADE} : "+fmt.Sprintf("%f", e.AmbienteC.UmidadeCorrenteValor()))
				produtorc.morrer()
			}

		}

	}

}

func (e *Ecossistema) FatoresAbioticosDeReproducaoProdutores (produtorc *Produtor) {

	if e.AmbienteC.LuzCorrenteValor() > produtorc._minLuzIdeal &&  e.AmbienteC.LuzCorrenteValor() < produtorc._maxLuzIdeal {
		produtorc.reproduzir()
		produtorc.reproduzir()
		fmt.Println("===> PRODUTOR :" + produtorc.toString() + "reproduziu por condições abioticas favoráveis")

	}
}

func randomBool() bool{
	return rand.Int() % 2 == 0
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

func (e *Ecossistema) GerarOrganismos(tipo string, quantidade int, nome string, adulto int, reproducao int, gestacao int, vida int, cor uint32, alimentacaoNome []string, nivelconsumidor int, comportamento string, temperaturaMin float32, temperaturaMax float32, umidadeMin float32, umidadeMax float32, ventoMax float32, morrePorChuvaEspecial string, minLuzIdeal float32, maxLuzIdeal float32) {

	switch tipo {

	case "produtor":
		for i := 0; i < quantidade; i++ {
			e.AdicionarProdutor(ProdutorNovo(nome, adulto, reproducao, gestacao, vida, cor, temperaturaMin, temperaturaMax, umidadeMin, umidadeMax, morrePorChuvaEspecial, minLuzIdeal, maxLuzIdeal, e))
		}

	case "consumidor":
		for i := 0; i < quantidade; i++ {
			e.AdicionarConsumidor(ConsumidorNovo(nome, adulto, reproducao, vida, cor, e, alimentacaoNome, nivelconsumidor, comportamento, temperaturaMin, temperaturaMax, ventoMax, morrePorChuvaEspecial))
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


			e.GerarOrganismos("produtor", 30, OrganismoNome, organismoC.Base.Adulto, organismoC.Reproducao.Frequencia,
				organismoC.Reproducao.Gestacao, organismoC.Base.Vida, cor, organismoC.Alimentacao.Nome, organismoC.Base.Nivel,
				organismoC.Sobrevivencia.Comportamento, organismoC.Sobrevivencia.TemperaturaMin, organismoC.Sobrevivencia.TemperaturaMax,
				organismoC.Sobrevivencia.UmidadeMin, organismoC.Sobrevivencia.UmidadeMax, organismoC.Sobrevivencia.VentoMax,
				organismoC.Sobrevivencia.MorrePorChuvaEspecial, organismoC.Sobrevivencia.MinLuzIdeal, organismoC.Sobrevivencia.MaxLuzIdeal)
		}
		if organismoC.Base.Tipo == "Consumidor" {
			var cor uint32 = organismoC.Base.Cor
			e.GerarOrganismos("consumidor", 10, OrganismoNome, organismoC.Base.Adulto, organismoC.Reproducao.Frequencia,
				organismoC.Reproducao.Gestacao, organismoC.Base.Vida, cor, organismoC.Alimentacao.Nome, organismoC.Base.Nivel,
				organismoC.Sobrevivencia.Comportamento, organismoC.Sobrevivencia.TemperaturaMin, organismoC.Sobrevivencia.TemperaturaMax,
				organismoC.Sobrevivencia.UmidadeMin, organismoC.Sobrevivencia.UmidadeMax, organismoC.Sobrevivencia.VentoMax, organismoC.Sobrevivencia.MorrePorChuvaEspecial, organismoC.Sobrevivencia.MinLuzIdeal, organismoC.Sobrevivencia.MaxLuzIdeal)
		}
	}

}

func (e *Ecossistema) Produtores_Contagem() int { return e.produtores_contagem }
func (e *Ecossistema) Produtores_Mortos() int   { return e.produtores_mortos }
func (e *Ecossistema) Produtores_Vivos() int    { return e.produtores_contagem - e.produtores_mortos }

func (e *Ecossistema) Consumidores_Contagem() int { return e.consumidores_contagem }
func (e *Ecossistema) Consumidores_Mortos() int   { return e.consumidores_mortos }
func (e *Ecossistema) Consumidores_Vivos() int    { return e.consumidores_contagem - e.consumidores_mortos }
