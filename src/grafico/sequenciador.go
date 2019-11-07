package grafico

// AUTOR : LUAN ALVES FREITAS
// DATA : 2019 09 02

type Sequenciador struct {
	valor1  float32
	valor2  float32
	valor3  float32
	valor4  float32
	valor5  float32
	valor6  float32
	valor7  float32
	valor8  float32
	valor9  float32
	valor10 float32
	valor11 float32
	valor12 float32
	valor13 float32
	valor14 float32
	valor15 float32
	valor16 float32
	valor17 float32
	valor18 float32
	valor19 float32
	valor20 float32

	i int
}

func SequenciaNova() *Sequenciador {
	p := Sequenciador{}

	p.i = 0

	return &p
}

func (p *Sequenciador) Adicionar(valor float32) {

	p.valor20 = p.valor19
	p.valor19 = p.valor18
	p.valor18 = p.valor17
	p.valor17 = p.valor16
	p.valor16 = p.valor15
	p.valor15 = p.valor14
	p.valor14 = p.valor13

	p.valor13 = p.valor12
	p.valor12 = p.valor11
	p.valor11 = p.valor10
	p.valor10 = p.valor9
	p.valor9 = p.valor8
	p.valor8 = p.valor7
	p.valor7 = p.valor6
	p.valor6 = p.valor5
	p.valor5 = p.valor4
	p.valor4 = p.valor3
	p.valor3 = p.valor2
	p.valor2 = p.valor1
	p.valor1 = valor

}

func (p *Sequenciador) ValorCorrente() float32 {
	var ret float32 = 0
	if p.i == 0 {
		ret = p.valor1
	} else if p.i == 1 {
		ret = p.valor2
	} else if p.i == 2 {
		ret = p.valor3
	} else if p.i == 3 {
		ret = p.valor4
	} else if p.i == 4 {
		ret = p.valor5
	} else if p.i == 5 {
		ret = p.valor6
	} else if p.i == 6 {
		ret = p.valor7
	} else if p.i == 7 {
		ret = p.valor8
	} else if p.i == 8 {
		ret = p.valor9
	} else if p.i == 9 {
		ret = p.valor10
	} else if p.i == 10 {
		ret = p.valor11
	} else if p.i == 11 {
		ret = p.valor12
	} else if p.i == 12 {
		ret = p.valor13
	} else if p.i == 13 {
		ret = p.valor14
	} else if p.i == 14 {
		ret = p.valor15
	} else if p.i == 15 {
		ret = p.valor16
	} else if p.i == 16 {
		ret = p.valor17
	} else if p.i == 17 {
		ret = p.valor18
	} else if p.i == 18 {
		ret = p.valor19
	} else if p.i == 19 {
		ret = p.valor20
	}

	p.i++
	if p.i >= 20 {
		p.i = 0
	}
	return ret
}
