package ecossistema

import (
	"math/rand"
	"strconv"
	"utils"
)

type umidificador struct {
	ambienteC *Ambiente

	umidade         float32
	umidadecontador int
	umidadelimite   int
	umidadefinalizador int


}

func umidificadorNovo(a *Ambiente) *umidificador {
	ret := umidificador{}
	ret.ambienteC = a

	ret.umidade = 0
	ret.umidadecontador = 16
	ret.umidadelimite = 15
ret.umidadefinalizador=ret.umidadelimite
	return &ret
}

func (a *umidificador) umidadeNomeCorrente() string {
	return a.umidadeNome(a.umidade)
}

func (a *umidificador) umidadeNome(_umidade float32) string {
	var ret string = ""

	if _umidade >= 0 && _umidade < 20 {
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

func (a *umidificador) umidificar() {

	if a.umidadecontador >= a.umidadefinalizador {
		a.umidadecontador = 0
		a.umidade = float32(rand.Intn(int(99))) + rand.Float32()

		a.umidadefinalizador=a.umidadelimite + rand.Intn(30)

		utils.Log("ambientelimitador.txt", "Umidade - "+strconv.Itoa(a.umidadefinalizador) + "   " + a.umidadeNomeCorrente())


	} else {

		if rand.Intn(100) < 50 {
			a.umidade += rand.Float32()

		} else {
			a.umidade -= rand.Float32()

		}

		a.umidadecontador++
	}
}
