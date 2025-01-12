package main

import "fmt"

//Ponteiros: é uma variável que ao invés de armazenar um valor,
//armazena um endereço de memória
//a memória tem um endereço e esse endereço que guarda o valor

func main() {
	i := 1
	fmt.Println("Valor inicial da variável i:", i)
	fmt.Println("Valor do end de memória do i:", &i)

	a := &i
	fmt.Println("A variável a recebe o end de mem da variável i:", a)
	fmt.Println("A variável *a mostra:", *a)
	fmt.Println("Variável a tem seu próprio endereço &a:", &a)
}
