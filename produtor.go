package main

import (
	"fmt"
)

type produtor struct {

  _nome        string
  _idade  int
  _status     string
  _fase       string

	adultociclo int
}

func Produtor_Novo(nome string, adulto int) *produtor {

  p := produtor{_nome: nome}
	p._idade = 0
	p._status = "vivo"
	p._fase = "nascido"
	p.adultociclo = adulto
	p._status = "vivo"
	p._fase = "nascido"

	return &p
}

func (p *produtor) vivendo() {

	if p._status == "vivo" {

	p._idade += 1

		if p._fase == "nascido" {
			if p._idade >= p.adultociclo {
				p._fase = "adulto"

        fmt.Println("       --- Produtor : ",p.nome() , " Evoluiu : Adulto !!!")

			}
		}

	}

}



func (p *produtor) nome() string   { return p._nome }
func (p *produtor) fase() string   { return p._fase }
func (p *produtor) status() string { return p._status }
func (p *produtor) ciclos() int   { return p._idade }
