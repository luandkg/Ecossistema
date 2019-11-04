package bioxml

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
)

func ListarOrganismos(local string) []string {

	var conteudo []string

	files, err := ioutil.ReadDir(local)
	for _, file := range files {

		if strings.HasSuffix(file.Name(), ".organismo") {

			var nomearquivo string = strings.Replace(file.Name(), ".organismo", "", 1)
			conteudo = append(conteudo, nomearquivo)
		}
	}
	err = err

	return conteudo
}

func CarregarOrganismo(local string) *Organismo {

	data, _ := ioutil.ReadFile(local)
	organismoC := &Organismo{}
	_ = xml.Unmarshal([]byte(data), &organismoC)

	return organismoC
}
