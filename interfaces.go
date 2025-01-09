package main

import "fmt"

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Cor   string
	Patas int
}

type Pessoa struct {
	Nome             string
	Idade            int
	Prof             string
	TempoDeProfissao int
}

func (p Pessoa) FalaBomDia() string {
	return fmt.Sprintf("%s deseja bom dia para você!", p.Nome)
}

func (p Pessoa) Profissao() string {
	return fmt.Sprintf("A profissão de %s é %s com %d anos de experiência", p.Nome, p.Prof, p.TempoDeProfissao)
}

// método atrelado a struct Cachorro
func (c Cachorro) Barulho() string {
	return "au-au"
}

// método atrelado a struct Gato
func (g Gato) Barulho() string {
	return "miau"
}

func (c Cachorro) NumDePatas() int {
	return c.Patas
}

func (g Gato) NumDePatas() int {
	return g.Patas
}

type Animal interface { //interface com dois métodos()
	Barulho() string //retorna string
	NumDePatas() int
}

// para entender que é humano, deve ter os dois métodos
type Adulto interface { //interface com dois métodos()
	FalaBomDia() string //retorna string
	Profissao() string
}

// função que recebe a interface animal diretamente, sem struct
func FazBarulho(animal Animal) {
	fmt.Println(animal.Barulho())
}

func main() {
	cachorro := Cachorro{
		Raca:  "Pischer",
		Cor:   "Preto e Marrom",
		Patas: 4,
	}

	gato := Gato{
		Cor:   "Cinza e Branco",
		Patas: 4,
	}

	fmt.Println(cachorro.Barulho(), cachorro.NumDePatas())
	fmt.Println(gato.Barulho(), gato.NumDePatas())

	FazBarulho(gato)
	FazBarulho(cachorro)

}
