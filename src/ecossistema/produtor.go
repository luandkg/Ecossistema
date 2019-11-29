package ecossistema

import (
	"fmt"
	"strconv"
	"utils"
)

type Produtor struct {
	organismo

	_adultociclo        int
	_reproduzirciclo    int
	_reproduzircontador int
	_reproduzirGestacao int

	_reproduzirEstaEmGestacao   bool
	_reproduzirGestacaoContador int

	_vida int

	_produtores   (map[string]*Produtor)
	_ecossistemaC *Ecossistema

	/*Variáveis para definição de Sobrevivência do Organismo*/
	_temperaturaMin float32
	_temperaturaMax float32
	_umidadeMin float32
	_umidadeMax float32
	_minLuzIdeal float32
	_maxLuzIdeal float32
	_morrePorChuvaEspecial string

}

// ProdutorNovo : Criar instancia de Produtor
func ProdutorNovo(nome string, adulto int, reproducao int, gestacao int, vida int, cor uint32, temperaturaMin float32, temperaturaMax float32, UmidadeMin float32, UmidadeMax float32, morrePorChuvaEspecial string, minLuzIdeal float32, maxLuzIdeal float32, ecossistemaC *Ecossistema) *Produtor {

	p := Produtor{_adultociclo: adulto}

	p.organismo = *OrganismoNovo(nome)
	p._nome = nome
	p._nomecompleto = ""
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

	p._reproduzirciclo = reproducao
	p._reproduzircontador = 0
	p._reproduzirGestacao = gestacao
	p._reproduzirEstaEmGestacao = false
	p._reproduzirGestacaoContador = 0

	p._vida = vida
	p._posx = 0
	p._posy = 0

	//	p.energizar(float32(adulto) * 12)
	p._energia = 0

	p._cor = cor

	/*Variáveis para definição de Sobrevivência do Organismo*/
	p._umidadeMin = UmidadeMin
	p._umidadeMax = UmidadeMax
	p._temperaturaMin = temperaturaMin
	p._temperaturaMax = temperaturaMax
	p._minLuzIdeal = minLuzIdeal
	p._maxLuzIdeal = maxLuzIdeal
	p._morrePorChuvaEspecial = morrePorChuvaEspecial

	p._ecossistemaC = ecossistemaC

	return &p
}

func (p *Produtor) vivendo() {

	p.organismo.vivendo()

	if p._status == "vivo" {

		if p._reproduzirEstaEmGestacao == true {

		} else {

		}

		if p._ecossistemaC.AmbienteC.fase == "Dia" {

			p._ecossistemaC.ProduzirOxigenio(+5)
			p._ecossistemaC.ProduzirCarbono(-5)

			var te float32 = ((p._ecossistemaC.AmbienteC.luz / 100) / 100)
			p.energizar(te)

		} else {
			p._ecossistemaC.ProduzirOxigenio(-5)
			p._ecossistemaC.ProduzirCarbono(+5)

		}

		if p._fase == "nascido" {
			p.jovem()
		}

		// Se o organismo for adulto inicia o ciclo de reproducao
		if p._fase == "adulto" {
			p.reproduzir()
		}

		if p._idade >= p._vida {
			p.morrer()
		}

	}

}

func (p *Produtor) jovem() {

	if p._idade >= p._adultociclo {
		p._fase = "adulto"

		fmt.Println("       --- Produtor : ", p.NomeCompleto(), " Evoluiu : Adulto !!!")

	}
}

func (p *Produtor) morrer() {

	p._status = "morto"
	fmt.Println("       --- Produtor : ", p.NomeCompleto(), " Morreu !!!")

}
func (p *Produtor) reproduzirLapsoTemporal() {

	//fmt.Println("       --- Produtor : ", p.Nome(), " ReproduzirContador : ",p._reproduzircontador)

	p._reproduzircontador += 1

	if p._reproduzircontador >= p._reproduzirciclo {
		p._reproduzircontador = 0
		p._reproduzirEstaEmGestacao = true
		p._reproduzirGestacaoContador = 0

		fmt.Println("       --- Produtor : ", p.NomeCompleto(), " Em Gestação !!!")

	}

}

func (p *Produtor) reproduzirEmGestacao() {

	p._reproduzirGestacaoContador++

	//fmt.Println("       --- Produtor : ", p.Nome(), " GestacaoContador : ",p._reproduzirGestacaoContador)

	if p._reproduzirGestacaoContador >= p._reproduzirGestacao {
		p._reproduzirEstaEmGestacao = false
		p._reproduzircontador = 0

		fmt.Println("       --- Produtor : ", p.NomeCompleto(), " Reproduzindo !!!")

		var pg = ProdutorNovo(p._nome, p._adultociclo, p._reproduzirciclo, p._reproduzirGestacao, p._vida, p._cor, p._temperaturaMin, p._temperaturaMax, p._umidadeMin, p._umidadeMax, p._morrePorChuvaEspecial, p._minLuzIdeal, p._maxLuzIdeal, p._ecossistemaC)
		var x int = utils.Aleatorionumero(50)
		var y int = utils.Aleatorionumero(50)

		pg.mudarposicao(x, y)

		p._ecossistemaC.AdicionarProdutor(pg)

	}

}

func (p *Produtor) reproduzir() {

	if p._reproduzirEstaEmGestacao == false {

		p.reproduzirLapsoTemporal()

	} else {

		p.reproduzirEmGestacao()

	}

}

func (p *Produtor) toString() string {

	s1 := " [" + p.Fase() + " " + strconv.Itoa(p.Ciclos()) + "]"
	s2 := " POS [" + strconv.Itoa(p.x()) + " " + strconv.Itoa(p.y()) + "]"
	s3 := " - Status : " + p._status
	s4 := "   -> { ENERGIA : " + fmt.Sprintf("%f", p._energia) + "}"

	return p.Nome() + s1 + s2 + s3 + s4 //+ strconv.FormatUint(p._cor, 10)
}
