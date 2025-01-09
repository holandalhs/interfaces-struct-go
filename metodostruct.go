package main

import "fmt"

// type <nome da estrutura> struct { <campos> }
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	AnoNasc   int
	Ativa     bool
}

func ListaNomeIdade(nome string, idade int) string {
	return fmt.Sprintf("%s tem %d anos", nome, idade)
}

func main() {
	//criar variável pessoa do tipo Pessoa
	pessoa := Pessoa{
		Nome:  "Mario",
		Idade: 30,
		Ativa: true,
	}

	println(ListaNomeIdade(pessoa.Nome, pessoa.Idade))

}
