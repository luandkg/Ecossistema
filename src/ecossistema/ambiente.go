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
	a.nublar()
	a.umidificar()




	if a.logciclo>=10{
a.logciclo=0
		utils.Log("ambiente.txt", "------------------------------------------------")

		utils.Log("ambiente.txt", "Dia  : " + strconv.Itoa(a.dia) + " FASE : " + a.fase + " CICLO : " + strconv.Itoa(a.ciclo) )


		s1 := fmt.Sprintf("%.2f ºC", a.temperaturaCorrente)
		//s2 := fmt.Sprintf("%f", a.chuva())
		//s3:= fmt.Sprintf("%f", a.nuvem)

		utils.Log("ambiente.txt", "Temperatura - " + s1)
		utils.Log("ambiente.txt", "Luz - " + a.luminosidadeNomeCorrente())
		utils.Log("ambiente.txt", "Nuvem - " + a.nuvemNomeCorrente())
		utils.Log("ambiente.txt", "Sol - " +  strconv.Itoa(a.Sol()))
		utils.Log("ambiente.txt", "Umidade - " + a.umidadeNomeCorrente())

		if a.ventorodando == true {
			utils.Log("ambiente.txt", "Vento - " + a.ventoCorrenteNome() + " [ " + a.ventoorigem + " -> " + a.ventodestino + " ] - " + a.ventomodo )
		}else{
			utils.Log("ambiente.txt", "Vento - " + a.ventoCorrenteNome() + " [ " + a.ventoorigem + " -> " + a.ventodestino + " ]")
		}

		utils.Log("ambiente.txt", "Chuva - " + a.chuvaNomeCorrente())


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
			fmt.Printf("\n\t VENTO %.2f %s [ %s : %s ] - Rodando : %s", a.vento, a.ventoCorrenteNome(), a.ventoorigem, a.ventodestino,a.ventomodo)
		} else {
			fmt.Printf("\n\t VENTO %.2f %s [ %s : %s ]", a.vento, a.ventoCorrenteNome(), a.ventoorigem, a.ventodestino)
		}


		fmt.Printf("\n\t Chuva :  %.2f ", a.chuva())


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
func (a *Ambiente) Sol() int { return int(a.luz - (a.nuvem/3))  }

func (a *Ambiente) chuva() float32 {
	var valor float32=0

	var fator1 float32 = a.vento/5
	var fator2 float32 = a.umidade/5
	var fator3 float32 = a.temperaturaCorrente
	var fator4 float32 = a.nuvem/5


	if a.temperaturaCorrente<=0 {
		fator3 = 25
	}else if a.temperaturaCorrente >=0 && a.temperaturaCorrente <20 {
			fator3=20
	}else if a.temperaturaCorrente >=20 && a.temperaturaCorrente <30 {
		fator3=10
	}else{
		fator3=0
	}

	valor = fator1+fator2+fator3+fator4

return valor
}

func (a *Ambiente) chuvaNomeCorrente() string {
	return a.chuvaNome(a.chuva())
}
func (a *Ambiente) chuvaNome(_chuva float32) string {
	var ret string = ""

	if _chuva >= 0 && _chuva < 40 {
		ret = "Sem Chuva"
	}

	if _chuva >= 40 && _chuva < 50 {
		ret = "Neblina"
	}

	if _chuva >= 50 && _chuva < 60 {
		ret = "Chuvisco"
	}

	if _chuva >= 60 && _chuva < 70 {
		ret = "Chuva"
	}

	if _chuva >= 80 {
		ret = "Chuva Forte"
	}

	return ret
}

func (a *Ambiente) ceu() string {
	return a.luminosidadeNomeCorrente()
}
