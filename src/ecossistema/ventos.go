package ecossistema

import "math/rand"

type ventos struct {
	ambienteC *Ambiente

	vento         float32
	ventocontador int
	ventolimite   int
	ventoorigem   string
	ventodestino  string
	ventorodando  bool
}

func ventosNovo(a *Ambiente) *ventos {
	ret := ventos{}
	ret.ambienteC = a

	ret.vento = 0
	ret.ventocontador = 16
	ret.ventolimite = 15
	ret.ventorodando = false

	return &ret
}

func (a *ventos) ventoCorrenteNome() string {
	return a.ventosNome(a.vento)
}

func (a *ventos) ventosNome(_vento float32) string {
	var ret string = ""

	if _vento >= 0 && _vento < 10 {
		ret = "Sem Vento"
	}

	if _vento >= 10 && _vento < 20 {
		ret = "Vento Leve"
	}

	if _vento >= 20 && _vento < 60 {
		ret = "Ventania"
	}

	if _vento >= 60 && _vento < 80 {
		ret = "Vento Fortes"
	}

	if _vento >= 80 {
		ret = "Vento Muito Forte"
	}

	return ret
}

func (a *ventos) ventar() {

	if a.ventocontador >= a.ventolimite {
		a.ventorodando = false
		a.ventocontador = 0
		a.vento = float32(rand.Intn(int(100))) + rand.Float32()

		var mudardirecao = rand.Intn(70)
		if mudardirecao >= 0 && mudardirecao < 10 {
			a.ventoorigem = "Norte"
		} else if mudardirecao >= 10 && mudardirecao < 20 {
			a.ventoorigem = "Leste"
		} else if mudardirecao >= 10 && mudardirecao < 20 {
			a.ventoorigem = "Sul"
		} else if mudardirecao >= 20 && mudardirecao < 30 {
			a.ventoorigem = "Oeste"
		} else if mudardirecao >= 30 && mudardirecao < 40 {
			a.ventoorigem = "Nordeste"
		} else if mudardirecao >= 40 && mudardirecao < 50 {
			a.ventoorigem = "Sudeste"
		} else if mudardirecao >= 50 && mudardirecao < 60 {
			a.ventoorigem = "Sudoeste"
		} else if mudardirecao >= 60 && mudardirecao < 70 {
			a.ventoorigem = "Noroeste"
		}

		mudardirecao = rand.Intn(70)
		if mudardirecao >= 0 && mudardirecao < 10 {
			a.ventodestino = "Norte"
		} else if mudardirecao >= 10 && mudardirecao < 20 {
			a.ventodestino = "Leste"
		} else if mudardirecao >= 10 && mudardirecao < 20 {
			a.ventodestino = "Sul"
		} else if mudardirecao >= 20 && mudardirecao < 30 {
			a.ventodestino = "Oeste"
		} else if mudardirecao >= 30 && mudardirecao < 40 {
			a.ventodestino = "Nordeste"
		} else if mudardirecao >= 40 && mudardirecao < 50 {
			a.ventodestino = "Sudeste"
		} else if mudardirecao >= 50 && mudardirecao < 60 {
			a.ventodestino = "Sudoeste"
		} else if mudardirecao >= 60 && mudardirecao < 70 {
			a.ventodestino = "Noroeste"
		}

		if a.ventoorigem == a.ventodestino {
			a.ventorodando = true
		}

	} else {

		if rand.Intn(100) < 50 {
			a.vento += rand.Float32()

		} else {
			a.vento -= rand.Float32()

		}

		a.ventocontador++
	}
}
