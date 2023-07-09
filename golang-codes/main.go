package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salário          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")

}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouce só: Bom dia!")

}

func main() {
	fmt.Println("Hello!")
}
