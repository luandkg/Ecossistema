package ecossistema

import (
	"fmt"
	"grafico"
	"strconv"

	"utils"
)

type Ambiente struct {
	fase         string
	faseciclo    int
	fasecontador int
	dia          int

	ciclo    int
	logciclo int

	gasOxigenio  float64
	gasCarbonico float64

	Temperatura
	Luminosidade
	Ventos
	Nuvens
	Umidificador
	Chuva
	Ceu
	Sensacao

	SeqTemperatura grafico.Sequenciador
	SeqUmidade     grafico.Sequenciador
	SeqVento       grafico.Sequenciador
	SeqNuvem       grafico.Sequenciador
	SeqLuz         grafico.Sequenciador
	SeqChuva       grafico.Sequenciador
}

func AmbienteNovo() *Ambiente {

	p := Ambiente{}
	p.faseciclo = 100
	p.dia = 1
	p.fase = "Dia"
	p.fasecontador = -1

	p.gasCarbonico = 0
	p.gasOxigenio = 0

	p.Temperatura = *TemperaturaNovo(&p)
	p.Luminosidade = *LuminosidadeNovo(&p)
	p.Ventos = *VentosNovo(&p)
	p.Nuvens = *NuvensNovo(&p)
	p.Umidificador = *UmidificadorNovo(&p)
	p.Chuva = *ChuvaNovo(&p)
	p.Ceu = *CeuNovo(&p)
	p.Sensacao = *SensacaoNovo(&p)

	p.ciclo = 0
	p.logciclo = 0
	return &p
}

func (a *Ambiente) AmbienteFase() {

	a.ProximoCiclo()

	a.Esquentar()
	a.Iluminar()
	a.Ventar()
	a.Nublar()
	a.Umidificar()
	a.Chover()

	a.log()

	a.SeqTemperatura.Adicionar(a.TemperaturaCorrente())
	a.SeqUmidade.Adicionar(a.UmidadeCorrenteValor())
	a.SeqVento.Adicionar(a.VentoCorrenteValor())
	a.SeqNuvem.Adicionar(a.NuvemCorrenteValor())
	a.SeqLuz.Adicionar(a.LuzCorrenteValor())
	a.SeqChuva.Adicionar(a.ChuvaCorrenteValor())

}

func (a *Ambiente) ProximoCiclo() {

	// Implementacao FASE - DIA / NOITE

	a.ciclo++

	if a.fasecontador >= a.faseciclo {
		a.fasecontador = 0
		if a.fase == "Dia" {
			a.fase = "Noite"
			//	utils.Log("logs/logs.txt", "Noite - "+strconv.Itoa(a.dia)+" [ ]")
		} else {
			a.fase = "Dia"
			a.dia++
			//utils.Log("logs/logs.txt", "Dia - "+strconv.Itoa(a.dia)+" [ "+a.CeuNomeCorrente()+"]")
		}
	} else {
		a.fasecontador++
	}

}

func (a *Ambiente) Fase() string      { return a.fase }
func (a *Ambiente) FaseContador() int { return a.fasecontador }
func (a *Ambiente) Ciclo() int        { return a.ciclo }
func (a *Ambiente) Dia() int          { return a.dia }

func (a *Ambiente) DiaInfo() string {
	return "Dia  : " + strconv.Itoa(a.dia) + " - [ FASE : " + a.fase + " CICLO : " + strconv.Itoa(a.ciclo) + " ] "
}

func (a *Ambiente) log() {

	a.logciclo++

	if a.logciclo >= 10 {
		a.logciclo = 0

		var identacao string = "     - "
		utils.Log("logs/ambiente.txt", "------------------------------------------------------------------------------------")
		utils.Log("logs/ambiente.txt", a.DiaInfo())
		utils.Log("logs/ambiente.txt", identacao+a.TemperaturaInfo())
		utils.Log("logs/ambiente.txt", identacao+a.LuzInfo())
		utils.Log("logs/ambiente.txt", identacao+a.NuvemInfo())
		utils.Log("logs/ambiente.txt", identacao+a.CeuInfo())
		utils.Log("logs/ambiente.txt", identacao+a.UmidadeInfo())
		utils.Log("logs/ambiente.txt", identacao+a.VentoInfo())
		utils.Log("logs/ambiente.txt", identacao+a.ChuvaInfo())
		utils.Log("logs/ambiente.txt", identacao+a.SensacaoInfo())

		fmt.Println("")
		fmt.Printf("\n\t " + a.DiaInfo())
		fmt.Println("")
		fmt.Printf("\n\t " + a.TemperaturaInfo())
		fmt.Printf("\n\t " + a.LuzInfo())
		fmt.Printf("\n\t " + a.NuvemInfo())
		fmt.Printf("\n\t " + a.CeuInfo())
		fmt.Printf("\n\t " + a.UmidadeInfo())
		fmt.Printf("\n\t " + a.VentoInfo())
		fmt.Printf("\n\t " + a.ChuvaInfo())
		fmt.Println()
		fmt.Println()
		fmt.Println()

		//utils.Log("logs/vento.txt", fmt.Sprintf("%.2f", a.vento))
		//utils.Log("logs/umidade.txt", fmt.Sprintf("%.2f", a.umidade))
		//utils.Log("logs/chuva.txt", fmt.Sprintf("%.2f", a.chover()))
		//utils.Log("logs/nuvem.txt", fmt.Sprintf("%.2f", a.nuvem))
		//utils.Log("logs/temperatura.txt", fmt.Sprintf("%.2f", a.TemperaturaCorrente))

	}

}