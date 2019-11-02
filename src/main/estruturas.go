package main

type Organismo struct {
	Base Base `xml:"Base"`
	Reproducao Reproducao `xml:"Reproducao"`
}

type Base struct {
	Tipo string `xml:"Tipo,attr"`
	Jovem        int `xml:"Jovem,attr"`
	Adulto       int `xml:"Adulto,attr"`
	Cor       uint32 `xml:"Cor,attr"`
}

type Reproducao struct {
	Frequencia        int `xml:"Frequencia,attr"`
}

