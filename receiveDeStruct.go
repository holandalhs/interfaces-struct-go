// Receive de Struct em GO, bastante utilizado em GO
package main

import "fmt"

// type <nome da estrutura> struct { <campos> }
type ReceivePessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	AnoNasc   int
	Ativa     bool
}

// atrelando o método ListaNomeEIdade a struct ReceivePessoa
// p - identificador da struct
func (p ReceivePessoa) ListaNomeEIdade() string { //retira os parâmetros
	return fmt.Sprintf("%s tem %d anos", p.Nome, p.Idade)
}

func main() {
	//criar variável pessoa do tipo Pessoa
	pessoa := ReceivePessoa{
		Nome:  "Mario",
		Idade: 30,
		Ativa: true,
	}

	println(pessoa.ListaNomeEIdade())

}
