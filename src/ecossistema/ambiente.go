package ecossistema

import (
	"fmt"
	"strconv"

	"utils"

	"github.com/veandco/go-sdl2/sdl"
)

type Ambiente struct {
	fase          string
	faseciclo     int
	fasecontador  int
	dia           int


	ciclo         int
	logciclo		int

	gasOxigenio float32
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

	p.gasCarbonico=0
	p.gasOxigenio=0

	p.temperatura = *temperaturaNovo(&p)
	p.luminosidade = *luminosidadeNovo(&p)
	p.ventos = *ventosNovo(&p)
	p.nuvens = *nuvensNovo(&p)
	p.umidificador = *umidificadorNovo(&p)

	p.ciclo = 0
p.logciclo=0
	return &p
}

func (a *Ambiente) ceu() string {

	if a.fase == "Dia" {
		return a.luminosidadeCorrenteNome()
	} else {
		return ""
	}

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
			utils.Log("logs.txt", "Noite - "+strconv.Itoa(a.dia)+" [ ]")

		} else {
			a.fase = "Dia"
			a.dia++

			utils.Log("logs.txt", "Dia - "+strconv.Itoa(a.dia)+" [ "+a.ceu()+"]")

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
	a.ceu()
a.umidificar()

	fmt.Println("")
	fmt.Println("Fase -> ", a.fase)
	fmt.Println("Quantidade de Sol -> ", a.luz)
	fmt.Println("Ceu -> ", a.ceu())

	fmt.Printf("\n\t Temperatura :  %.2f NG", a.temperaturaCorrente)

	if a.ventorodando == true {
		fmt.Printf("\n\t VENTO %.2f %s [ %s : %s ] - Rodando ", a.vento, a.ventoCorrenteNome(), a.ventoorigem, a.ventodestino)
	} else {
		fmt.Printf("\n\t VENTO %.2f %s [ %s : %s ]", a.vento, a.ventoCorrenteNome(), a.ventoorigem, a.ventodestino)
	}

	fmt.Printf("\n\t Luz :  %.2f - %s", a.luz, a.luminosidadeCorrenteNome())
	fmt.Printf("\n\t Nuvens :  %.2f - %s", a.nuvem, a.nuvemCorrente())

	if a.logciclo>=10{
a.logciclo=0
		utils.Log("ambiente.txt", "------------------------------------------------")

		utils.Log("ambiente.txt", "Dia  : " + strconv.Itoa(a.dia) + " FASE : " + a.fase + " CICLO : " + strconv.Itoa(a.ciclo) )


		s1 := fmt.Sprintf("%f", a.temperaturaCorrente)
		s2 := fmt.Sprintf("%f", a.chuva())

		utils.Log("ambiente.txt", "Temperatura - " + s1)
		utils.Log("ambiente.txt", "Luz - " + a.luminosidadeCorrenteNome())
		utils.Log("ambiente.txt", "Nuvem - " + a.nuvemCorrente())
		utils.Log("ambiente.txt", "Umidade - " + a.umidadeNomeCorrente())
		utils.Log("ambiente.txt", "Chuva - " + s2)

		if a.ventorodando == true {
			utils.Log("ambiente.txt", "Vento - " + a.ventoCorrenteNome() + " [ " + a.ventoorigem + " -> " + a.ventodestino + " ] - SIM ")
		}else{
			utils.Log("ambiente.txt", "Vento - " + a.ventoCorrenteNome() + " [ " + a.ventoorigem + " -> " + a.ventodestino + " ]")
		}



	}


	fmt.Println()

	fmt.Println()
}


func (a *Ambiente) AtualizarTela(s *sdl.Surface,) {

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

func (a *Ambiente) Fase() string { return a.fase }
func (a *Ambiente) FaseContador() int { return a.fasecontador }
func (a *Ambiente) Ciclo() int { return a.ciclo }
func (a *Ambiente) Sol() int { return int(a.luz) }

func (a *Ambiente) chuva() float32 {
	var valor float32=0

	if a.vento>=50 && a.vento<80{
		valor+=15
	}

	if a.vento>=80 && a.vento<100{
		valor+=30
	}

	if a.umidade>=50 && a.umidade<80{
		valor+=15
	}

	if a.umidade>=80 && a.umidade<100{
		valor+=30
	}

	if a.temperaturaCorrente<=33{
		valor+=15
	}

	if a.nuvem>=50{
		valor+=15
	}

return valor
}