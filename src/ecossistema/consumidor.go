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
	
	_nivelconsumidor	int
	_alimentacao		[]string

	_temAlvo            bool
	_alvoX              int
	_alvoY              int

	_vida int

	_ecossistemaC *Ecossistema
}

func ConsumidorNovo(nome string, adulto int, reproducao int, vida int, cor uint32, ecossistemaC *Ecossistema, alimentacaoNome []string, nivelconsumidor int) *Consumidor {

	p := Consumidor{_adultociclo: adulto}

	p.organismo = *OrganismoNovo(nome)
	p._nome = nome
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

	p._reproduzirciclo = reproducao
	p._reproduzircontador = 0

	p._nivelconsumidor = nivelconsumidor
	p._alimentacao = alimentacaoNome

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

		var pg = ConsumidorNovo(c._nome, c._adultociclo, c._reproduzirciclo, c._vida, c._cor, c._ecossistemaC, c._alimentacao, c._nivelconsumidor)
		var x = utils.Aleatorionumero(50)
		var y = utils.Aleatorionumero(50)

		peca := tb.RecuperarPeca(x, y)

		if !peca.VerificarPosicao() {
			pg.mudarposicao(x, y)
			peca.OcuparPosicao()
		}

		c._ecossistemaC.AdicionarConsumidor(pg)
	}

}

func verificaAlimento(nome string, alimentacao []string) bool{
	for _, a := range alimentacao {
        if a == nome {
            return true
        }
    }
    return false
}

func (c *Consumidor) VerificarAlvo(tb *tabuleiro.Tabuleiro) {

	if !c._temAlvo {

		peca := tb.RecuperarPeca(c._alvoX, c._alvoY)

		if !peca.VerificarPosicao() {

			var alvo = false
			var alvoX = 0
			var alvoY = 0

			var tetoBusca = 10
			var chaoBusca = -tetoBusca

			// TODO: Refatorar com posicoes absolutas no tabuleiro
			// TODO Refatorar e retirar for duplicado
			if (c._nivelconsumidor == 1){

				for _, produtor := range c._ecossistemaC.produtores {

					for i := tetoBusca; i > chaoBusca; i-- {

						for j := tetoBusca; j > chaoBusca; j-- {

							if produtor._posx == c._posx+i && produtor._posy == c._posy+j && verificaAlimento(produtor.organismo._nome, c._alimentacao){
								alvo = true
								alvoX = c._posx + i
								alvoY = c._posy + j
								break
							}
						}

						if alvo {
							break
						}

					}

					if alvo {
						break
					}

				}
			}else{
				for _, consumidor := range c._ecossistemaC.consumidores {

					for i := tetoBusca; i > chaoBusca; i-- {

						for j := tetoBusca; j > chaoBusca; j-- {

							if consumidor._posx == c._posx+i && consumidor._posy == c._posy+j && verificaAlimento(consumidor.organismo._nome, c._alimentacao){
								alvo = true
								alvoX = c._posx + i
								alvoY = c._posy + j
								break
							}
						}

						if alvo {
							break
						}

					}

					if alvo {
						break
					}

				}

			}

			c._temAlvo = alvo
			c._alvoX = alvoX
			c._alvoY = alvoY

		}

	}

}

func (c *Consumidor) reduzirDistanciaAlvo(tb *tabuleiro.Tabuleiro) {

	var novoEspacoEncontrado = false
	var contadorEspacoTentado = 0
	var distanciaX = c._posx - c._alvoX
	var distanciaY = c._posy - c._alvoY

	for !novoEspacoEncontrado {

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

		peca := tb.RecuperarPeca(novoX, novoY)

		if !peca.VerificarPosicao() {
			c.mudarposicao(novoX, novoY)
			peca.OcuparPosicao()
			novoEspacoEncontrado = true
		}

		if contadorEspacoTentado >= 3 {
			c.apagarAlvo()
			break
		} else {
			contadorEspacoTentado++
		}

	}

}

func (c *Consumidor) CacarAlvo(tb *tabuleiro.Tabuleiro) {

	var distanciaX = c._posx - c._alvoX
	var distanciaY = c._posy - c._alvoY

	if distanciaX == 1 || distanciaX == -1 {
		if distanciaY >= -1 && distanciaY <= 1 {
			c.matarAlvo()
			c.apagarAlvo()
		} else {
			c.reduzirDistanciaAlvo(tb)
		}
	} else if distanciaX == 0 {
		if distanciaY == 1 || distanciaY == -1 {
			c.matarAlvo()
			c.apagarAlvo()
		} else {
			c.reduzirDistanciaAlvo(tb)
		}
	} else {
		c.reduzirDistanciaAlvo(tb)
	}

}

func (c *Consumidor) matarAlvo() {

	// TODO: Refatorar com posicoes absolutas no tabuleiro
	if (c._nivelconsumidor == 1){
		for _, produtor := range c._ecossistemaC.produtores {

			if produtor._posx == c._alvoX && produtor._posy == c._alvoY {
				produtor.morrer()
				break
			}

		}
	}else{
		for _, consumidor := range c._ecossistemaC.consumidores {

			if consumidor._posx == c._alvoX && consumidor._posy == c._alvoY {
				consumidor.morrer()
				break
			}

		}
	}

}

func (c *Consumidor) apagarAlvo() {

	c._temAlvo = false
	c._alvoX = 0
	c._alvoY = 0

}

func (c *Consumidor) TemAlvo() bool { return c._temAlvo }

func (c *Consumidor) toString() string {

	var str = c.Nome() + " [" + c.Fase() + " " + strconv.Itoa(c.Ciclos()) + "]" + " POS[" + strconv.Itoa(c.x()) + " " + strconv.Itoa(c.y()) + "]"

	return str
}

//TODO inserido na parte de consumidor seucnario
func (p *Consumidor) morrer() {

	p._status = "morto"
	fmt.Println("       --- Consumidor : ", p.Nome(), " Morreu !!!")
}
