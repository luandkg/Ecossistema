package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type planta struct {
	organismo
	_adultociclo        int
	_reproduzirciclo    int
	_reproduzircontador int

	_vida int

	_direcao       string
	_dirquantidade int
	_dircontador   int
}

// Plantanovo : Criar instancia de planta
func Plantanovo(nome string, adulto int, reproducao int, vida int) *planta {

	p := planta{_adultociclo: adulto}
	p._nome = nome
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p._adultociclo = adulto

	p._reproduzirciclo = reproducao
	p._reproduzircontador = 0

	p._vida = vida
	p._posx = 0
	p._posy = 0

	return &p
}

func (p *planta) vivendo() {

	p.organismo.vivendo()

	if p._status == "vivo" {

		if p._fase == "nascido" {
			if p._idade >= p._adultociclo {
				p._fase = "adulto"

				fmt.Println("       --- Planta : ", p.nome(), " Evoluiu : Adulto !!!")

			}
		}

		// Se o organismo for adulto inicia o ciclo de reproducao
		if p._fase == "adulto" {

			p._reproduzircontador += 1

			if p._reproduzircontador >= p._reproduzirciclo {
				p._reproduzircontador = 0
				fmt.Println("       --- Planta : ", p.nome(), " Reproduzindo !!!")
			}

		}

		if p._idade >= p._vida {
			//p._status = "morto"
			fmt.Println("       --- Planta : ", p.nome(), " Morreu !!!")
		}

	}

}

func (p *planta) toString() string {

	var str = p.nome() + " [" + p.fase() + " " + strconv.Itoa(p.ciclos()) + "]" + " POS[" + strconv.Itoa(p.x()) + " " + strconv.Itoa(p.y()) + "]"

	return str
}

func (p *planta) movimento() {

	p._dircontador += 1

	if p._dircontador >= p._dirquantidade {
		p._dircontador = 0
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		p._dirquantidade = r1.Intn(15)

		p._direcao = ""
	}

	if p._direcao == "" {
		p._direcao = "l"

		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		escolha := r1.Intn(3)

		if escolha == 0 {
			p._direcao = "l"
		}

		if escolha == 1 {
			p._direcao = "o"
		}

		if escolha == 2 {
			p._direcao = "s"
		}

		if escolha == 3 {
			p._direcao = "n"
		}
	}

	if p._direcao == "l" {
		p._posx += 1
		if p._posx >= 50 {
			p._direcao = "o"
			p._posx = 48
		}
	} else if p._direcao == "o" {
		p._posx -= 1

		if p._posx < 0 {
			p._direcao = "l"
			p._posx = 1
		}

	} else if p._direcao == "n" {
		p._posy -= 1

		if p._posy < 0 {
			p._direcao = "s"
			p._posy = 1
		}
	} else if p._direcao == "s" {
		p._posy += 1

		if p._posy >= 50 {
			p._direcao = "n"
			p._posy = 48
		}

	}
}
