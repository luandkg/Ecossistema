package ecossistema

import (
	"fmt"
	"strconv"
	"tabuleiro"

	"utils"
)

type Consumidor struct {
	organismo
	_adultociclo        int
	_reproduzirciclo    int
	_reproduzircontador int

	_temAlvo            bool
	_alvoX              int
	_alvoY              int

	_vida int

	_ecossistemaC *Ecossistema
}

func ConsumidorNovo(nome string, adulto int, reproducao int, vida int, cor uint32, ecossistemaC *Ecossistema) *Consumidor {

	p := Consumidor{_adultociclo: adulto}

	p.organismo = *OrganismoNovo(nome)
	p._nome = nome
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

	p._reproduzirciclo = reproducao
	p._reproduzircontador = 0

	p._temAlvo = false
	p._alvoX = 0
	p._alvoY = 0

	p._vida = vida
	p._posx = 0
	p._posy = 0

	p._cor = cor
	p._ecossistemaC = ecossistemaC

	return &p

}

func (c *Consumidor) vivendo(tb *tabuleiro.Tabuleiro) {

	c.organismo.vivendo()

	if c._status == "vivo" {

		if c._idade == c._adultociclo || c._idade == c._vida {

			c.mudarFase()

		}

		if c._fase == "adulto" && c._idade < c._vida {

			c.reproduzir(tb)

		}

		if c._idade >= c._vida {
			c._status = "morto"
			fmt.Println("       --- Consumidor : ", c.Nome(), " Morreu !!!")
		}
	}

}

func (c *Consumidor) mudarFase() {

	switch c._fase {

	case "nascido":
		c._fase = "adulto"
		fmt.Println("       --- Consumidor : ", c.Nome(), " Evoluiu : Adulto !!!")

	case "adulto":
		c._status = "morto"
		c._fase = "falescido"
		fmt.Println("       --- Consumidor : ", c.Nome(), " Morreu !!!")

	}

}

func (c *Consumidor) reproduzir(tb *tabuleiro.Tabuleiro) {

	c._reproduzircontador += 1

	if c._reproduzircontador >= c._reproduzirciclo {
		c._reproduzircontador = 0
		fmt.Println("       --- Consumidor : ", c.Nome(), " Reproduzindo !!!")

		var pg = ConsumidorNovo(c._nome, c._adultociclo, c._reproduzirciclo, c._vida, c._cor, c._ecossistemaC)
		var x int = utils.Aleatorionumero(50)
		var y int = utils.Aleatorionumero(50)

		pg.mudarposicao(x, y)

		peca := tb.RecuperarPeca(x, y)

		if peca.VerificarPosicao() == false {
			pg.mudarposicao(x, y)
			peca.OcuparPosicao()
		}

		c._ecossistemaC.AdicionarConsumidor(pg)
	}

}

func (c *Consumidor) VerificarAlvo(p map[string]*Produtor) {

	var tetoBusca = 5
	var chaoBusca = -tetoBusca

	for _, produtor := range p {

		for i := tetoBusca; i >= chaoBusca; i-- {

			for j := tetoBusca; j >= chaoBusca; j-- {
				if produtor._posx == c._posx + i && produtor._posy == c._posy + j {
					c._temAlvo = true
					c._alvoX = c._posx + i
					c._alvoY = c._posy + j
					break
				}
			}

			if c._temAlvo {
				break
			}

		}

		if c._temAlvo {
			break
		}

	}

}

func (c *Consumidor) reduzirDistanciaAlvo() {

	var distanciaX = c._posx - c._alvoX
	var distanciaY = c._posy - c._alvoY

	var novoX = c._posx
	var novoY = c._posy

	if distanciaX < -1 {
		novoX += 1
	} else if distanciaX > 1 {
		novoX -= 1
	}

	if distanciaY < -1 {
		novoY += 1
	} else if distanciaY > 1 {
		novoY -= 1
	}

	c._posx = novoX
	c._posy = novoY

	fmt.Println("Novo X: ", c._posx, " Novo Y: ", c._posy, " Alvo X: ", c._alvoX, " Alvo Y: ", c._alvoY)

}

func (c *Consumidor) CacarAlvo() {

	var distanciaX = c._posx - c._alvoX
	var distanciaY = c._posy - c._alvoY

	fmt.Println("X: ", c._posx, " Y: ", c._posy, " Alvo X: ", c._alvoX, " Alvo Y: ", c._alvoY)
	fmt.Println("Distancia X: ", distanciaX, " Distancia Y: ", distanciaY)

	if distanciaX == 1 || distanciaY == -1 {
		if distanciaY > -1 && distanciaY < 1 {
			// matar planta
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!Atacar alvo")
		} else {
			c.reduzirDistanciaAlvo()
		}
	} else if distanciaY == 0 {
		if distanciaY == 1 || distanciaY == -1 {
			// matar planta
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!Atacar alvo")
		} else {
			c.reduzirDistanciaAlvo()
		}
	} else {
		c.reduzirDistanciaAlvo()
	}

}

func (c *Consumidor) TemAlvo() bool { return c._temAlvo }

func (c *Consumidor) toString() string {

	var str = c.Nome() + " [" + c.Fase() + " " + strconv.Itoa(c.Ciclos()) + "]" + " POS[" + strconv.Itoa(c.x()) + " " + strconv.Itoa(c.y()) + "]"

	return str
}
