package main

import (
"io/ioutil"
"strings"
)

type OrganismosXML struct {

}

func OrganismosXMLNovo() *OrganismosXML {
	ret := OrganismosXML{}
	return &ret
}

func (a*OrganismosXML) listar(local string) []string {

	var conteudo []string


	files, err := ioutil.ReadDir(local)
	for _, file := range files {

		if strings.HasSuffix(file.Name(), ".organismo") {

			var nomearquivo string=strings.Replace(file.Name(),".organismo","",1)
			conteudo = append(conteudo, nomearquivo)
		}
	}
	err=err

	return conteudo
}
