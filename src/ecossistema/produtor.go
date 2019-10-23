package ecossistema

import (
	"fmt"
	"strconv"
	"utils"
)

type Produtor struct {
	organismo

	_nomecompleto       string
	_adultociclo        int
	_reproduzirciclo    int
	_reproduzircontador int

	_vida int

	_produtores   (map[string]*Produtor)
	_ecossistemaC *Ecossistema
}

// ProdutorNovo : Criar instancia de Produtor
func ProdutorNovo(nome string, adulto int, reproducao int, vida int, cor uint32, ecossistemaC *Ecossistema) *Produtor {

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

	p._vida = vida
	p._posx = 0
	p._posy = 0

	p.energizar(float32(adulto)*12)

	p._cor = cor
	p._ecossistemaC = ecossistemaC

	return &p
}

func (p *Produtor) vivendo() {

	p.organismo.vivendo()

	if p._status == "vivo" {

		p.energizar(-p._ecossistemaC.ambienteC.luz/150)

		p._ecossistemaC.produzirOxigenio(0.0005)
		p._ecossistemaC.produzirCarbono(-0.00075)

		p.energizar(p._ecossistemaC.ambienteC.luz/100)

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

		fmt.Println("       --- Produtor : ", p.Nome(), " Evoluiu : Adulto !!!")

	}
}


func (p *Produtor) morrer() {

	p._status = "morto"
	fmt.Println("       --- Produtor : ", p.Nome(), " Morreu !!!")

}

func (p *Produtor) reproduzir() {

	p._reproduzircontador += 1

	if p._reproduzircontador >= p._reproduzirciclo {
		p._reproduzircontador = 0
		fmt.Println("       --- Produtor : ", p.Nome(), " Reproduzindo !!!")

		var pg = ProdutorNovo(p._nome, p._adultociclo, p._reproduzirciclo, p._vida, p._cor, p._ecossistemaC)
		var x int = utils.Aleatorionumero(50)
		var y int = utils.Aleatorionumero(50)

		pg.mudarposicao(x, y)

		p._ecossistemaC.AdicionarProdutor(pg)
	}

}

func (p *Produtor) toString() string {

	s1:=fmt.Sprintf("%f", p._energia)

	var str = p.Nome() + " [" + p.Fase() + " " + strconv.Itoa(p.Ciclos()) + "]" + " POS[" + strconv.Itoa(p.x()) + " " + strconv.Itoa(p.y()) + "] - Status : " + p._status + "   -> { ENERGIA : " + s1  + "}"

	return str
}
