package ecossistema

import (
	"fmt"
	"math/rand"
)

type Umidificador struct {
	ambienteC *Ambiente

	umidade            float32
	umidadecontador    int
	umidadelimite      int
	umidadefinalizador int
}

func UmidificadorNovo(a *Ambiente) *Umidificador {
	ret := Umidificador{}
	ret.ambienteC = a

	ret.umidade = 50.00
	ret.umidadecontador = 80
	ret.umidadelimite = 15
	ret.umidadefinalizador = ret.umidadelimite
	return &ret
}

func (a *Umidificador) UmidadeCorrenteValor() float32 {
	return a.umidade
}

func (a *Umidificador) umidadeNomeCorrente() string {
	return a.UmidadeNome(a.umidade)
}

func (a *Umidificador) UmidadeCorrente() string {
	return a.UmidadeNome(a.umidade)
}

func (a *Umidificador) ArSecoStatus() string {
	return a.ArSeco(a.umidade)
}

func (a *Umidificador) ArSeco(_umidade float32) string {

	var ret string = ""

	if _umidade < 12 {
		ret = "Emergência"
	}

	if _umidade >= 12 && _umidade <= 20 {
		ret = "Altera"
	}

	if _umidade >= 21 && _umidade <= 30 {
		ret = "Atenção"
	}

	if _umidade >= 31 && _umidade <= 40 {
		ret = "Observável"
	}

	if _umidade >= 41 {
		ret = "Não"
	}

	return ret

}



func (a *Umidificador) UmidadeNome(_umidade float32) string {
	var ret string = ""

	if  _umidade < 20 {
		ret = "Muito Baixa"
	}

	if _umidade >= 20 && _umidade < 40 {
		ret = "Baixa"
	}

	if _umidade >= 40 && _umidade < 60 {
		ret = "Normal"
	}

	if _umidade >= 60 && _umidade < 80 {
		ret = "Alta"
	}

	if _umidade >= 80 {
		ret = "Muito Alta"
	}

	return ret
}

func (a *Umidificador) Umidificar() {

	if a.umidadecontador >= a.umidadefinalizador {
		a.umidadecontador = 0
		a.umidade = float32(rand.Intn(int(99))) + rand.Float32()

		a.umidadefinalizador = a.umidadelimite + rand.Intn(30)

		//utils.Log("logs/ambientelimitador.txt", "Umidade - "+strconv.Itoa(a.umidadefinalizador) + "   " + a.umidadeNomeCorrente())

	} else {

		if rand.Intn(100) < 50 {
			a.umidade += rand.Float32()

		} else {
			a.umidade -= rand.Float32()

		}

		a.umidadecontador++
	}
}

func (a *Umidificador) UmidadeInfo() string {
	var ret string = fmt.Sprintf("Umidade :  %.2f - %s", a.UmidadeCorrenteValor(), a.UmidadeCorrente())

	return ret
}
