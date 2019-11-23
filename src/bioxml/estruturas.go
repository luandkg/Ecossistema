package bioxml

type Organismo struct {
	Base       Base       `xml:"Base"`
	Reproducao Reproducao `xml:"Reproducao"`
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

type Taxas struct {
	GasCarbonico int `xml:"GasCarbonico,attr"`
	GasOxigenio  int `xml:"GasOxigenio,attr"`
}

type Alimentacao struct{
	Nome []string `xml:"Nome"`
}
