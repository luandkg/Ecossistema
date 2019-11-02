package ecossistema

import (
	"fmt"
	"strconv"

	"utils"

	"github.com/veandco/go-sdl2/sdl"
)

type Ambiente struct {
	fase         string
	faseciclo    int
	fasecontador int
	dia          int

	ciclo    int
	logciclo int

	gasOxigenio  float32
	gasCarbonico float32

	temperatura
	luminosidade
	ventos
	nuvens
	umidificador
}

func AmbienteNovo() *Ambiente {

	p := Ambiente{}
	p.faseciclo = 100
	p.dia = 0
	p.fase = ""
	p.fasecontador = 0

	p.gasCarbonico = 0
	p.gasOxigenio = 0

	p.temperatura = *temperaturaNovo(&p)
	p.luminosidade = *luminosidadeNovo(&p)
	p.ventos = *ventosNovo(&p)
	p.nuvens = *nuvensNovo(&p)
	p.umidificador = *umidificadorNovo(&p)

	p.ciclo = 0
	p.logciclo = 0

	return &p
}

func (a *Ambiente) AmbienteFase() {

	// Implementacao FASE - DIA / NOITE

	a.ciclo++
	a.logciclo++

	if a.fase == "" {
		a.fasecontador = a.faseciclo * 2
	}

	if a.fasecontador >= a.faseciclo {
		a.fasecontador = 0
		if a.fase == "Dia" {
			a.fase = "Noite"
			utils.Log("logs/logs.txt", "Noite - "+strconv.Itoa(a.dia)+" [ ]")

		} else {
			a.fase = "Dia"
			a.dia++

			utils.Log("logs/logs.txt", "Dia - "+strconv.Itoa(a.dia)+" [ "+a.ceu()+"]")

		}
	} else {
		a.fasecontador++

		if a.fase == "Dia" {

		}
	}

	a.temperaturaDia()
	a.temperaturaNoite()

	a.claridade()
	a.ventar()
	a.nublar()
	a.umidificar()
	a.chover()

	if a.logciclo >= 10 {
		a.logciclo = 0

		utils.Log("ambiente.txt", "------------------------------------------------")

		utils.Log("ambiente.txt", "Dia  : "+strconv.Itoa(a.dia)+" FASE : "+a.fase+" CICLO : "+strconv.Itoa(a.ciclo))

		s1 := fmt.Sprintf("%.2f ºC", a.temperaturaCorrente)
		//s2 := fmt.Sprintf("%f", a.chuva())
		//s3:= fmt.Sprintf("%f", a.nuvem)

		utils.Log("logs/ambiente.txt", "Temperatura - "+s1)
		utils.Log("logs/ambiente.txt", "Luz - "+a.luminosidadeNomeCorrente())
		utils.Log("logs/ambiente.txt", "Nuvem - "+a.nuvemNomeCorrente())
		utils.Log("logs/ambiente.txt", "Sol - "+strconv.Itoa(a.Sol()))
		utils.Log("logs/ambiente.txt", "Umidade - "+a.umidadeNomeCorrente())

		if a.ventorodando == true {
			utils.Log("logs/ambiente.txt", "Vento - "+a.ventoCorrenteNome()+" [ "+a.ventoorigem+" -> "+a.ventodestino+" ] - "+a.ventomodo)
		} else {
			utils.Log("logs/ambiente.txt", "Vento - "+a.ventoCorrenteNome()+" [ "+a.ventoorigem+" -> "+a.ventodestino+" ]")
		}

		utils.Log("logs/ambiente.txt", "Chuva - "+a.chuvaNomeCorrente())

		fmt.Println("")
		fmt.Println("Fase -> ", a.fase)
		fmt.Println("Quantidade de Sol -> ", a.luz)
		fmt.Println("Ceu -> ", a.ceu())

		fmt.Printf("\n\t Temperatura :  %.2f ºC", a.temperaturaCorrente)
		fmt.Printf("\n\t Luz :  %.2f - %s", a.luz, a.luminosidadeNomeCorrente())
		fmt.Printf("\n\t Nuvens :  %.2f - %s", a.nuvem, a.nuvemNomeCorrente())
		fmt.Printf("\n\t Sol :  %.2f", a.Sol())
		fmt.Printf("\n\t Umidade :  %.2f - %s", a.umidade, a.umidadeNomeCorrente())

		if a.ventorodando == true {
			fmt.Printf("\n\t VENTO %.2f %s [ %s : %s ] - Rodando : %s", a.vento, a.ventoCorrenteNome(), a.ventoorigem, a.ventodestino, a.ventomodo)
		} else {
			fmt.Printf("\n\t VENTO %.2f %s [ %s : %s ]", a.vento, a.ventoCorrenteNome(), a.ventoorigem, a.ventodestino)
		}

		fmt.Printf("\n\t Chuva :  %.2f ", a.chover())

	}

	fmt.Println()

	fmt.Println()


	utils.Log("logs/vento.txt", fmt.Sprintf("%.2f", a.vento))
	utils.Log("logs/umidade.txt", fmt.Sprintf("%.2f", a.umidade))
	utils.Log("logs/chuva.txt", fmt.Sprintf("%.2f", a.chover()))
	utils.Log("logs/nuvem.txt", fmt.Sprintf("%.2f", a.nuvem))
	utils.Log("logs/temperatura.txt", fmt.Sprintf("%.2f", a.temperaturaCorrente))




}

func (a *Ambiente) AtualizarTela(s *sdl.Surface, ) {

	var linhafinal = sdl.Rect{0, 500, 500, 10}
	if a.Fase() == "Dia" {
		s.FillRect(&linhafinal, 0xFFFF00)
	} else {
		s.FillRect(&linhafinal, 0x000080)
	}

	var st = a.Sol() * 5
	var solinha = sdl.Rect{0, 510, int32(st), 10}
	s.FillRect(&solinha, 0xFF4500)

}

func (a *Ambiente) Fase() string      { return a.fase }
func (a *Ambiente) FaseContador() int { return a.fasecontador }
func (a *Ambiente) Ciclo() int        { return a.ciclo }
func (a *Ambiente) Sol() int          { return int(a.luz - (a.nuvem / 3)) }

func (a *Ambiente) ceu() string {
	return a.ceuLuminosidadeNome(a.Sol())
}
