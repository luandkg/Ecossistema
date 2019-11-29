package bioxml

type Organismo struct {
	Base       Base       `xml:"Base"`
	Reproducao Reproducao `xml:"Reproducao"`
	Sobrevivencia Sobrevivencia `xml:"Sobrevivencia"`
	Taxas      Taxas      `xml:"Taxas"`
	Alimentacao Alimentacao `xml:"Alimentacao"`
}

type Base struct {
	Tipo   string `xml:"Tipo,attr"`
	Nivel int `xml:"Nivel,attr"`
	Adulto int    `xml:"Adulto,attr"`
	Vida   int    `xml:"Vida,attr"`
	Cor    uint32 `xml:"Cor,attr"`
}

type Reproducao struct {
	Frequencia int `xml:"Frequencia,attr"`
	Gestacao   int `xml:"Gestacao,attr"`
}

type Sobrevivencia struct {
	Comportamento string `xml:"Comportamento,attr"`
	TemperaturaMin float32 `xml:"TemperaturaMin,attr"`
	TemperaturaMax float32 `xml:"TemperaturaMax,attr"`
	VentoMax float32 `xml:"VentoMax,attr"`
	MorrePorChuvaEspecial string `xml:"MorrePorChuvaEspecial,attr"`
	UmidadeMin float32 `xml:"UmidadeMin,attr"`
	UmidadeMax float32 `xml:"UmidadeMax,attr"`
	MinLuzIdeal float32 `xml:"minLuzIdeal,attr"`
	MaxLuzIdeal float32 `xml:"maxLuzIdeal,attr"`
}

type Taxas struct {
	GasCarbonico int `xml:"GasCarbonico,attr"`
	GasOxigenio  int `xml:"GasOxigenio,attr"`
}

type Alimentacao struct{
	Nome []string `xml:"Nome"`
}
