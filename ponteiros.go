package main

import "fmt"

//Ponteiros: é uma variável que ao invés de armazenar um valor,
//armazena um endereço de memória
//a memória tem um endereço e esse endereço que guarda o valor

func porcentValue(p int) {
	p = p * 100
}

func zeroValue(i int) {
	i = 0
	fmt.Println("End de memória do i dentro da func zeroValue:", &i)
	fmt.Println("Já o valor do i na func zeroValue é:", i)
}

func zeroPointer(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("Valor inicial da variável i:", i)
	fmt.Println("Valor do end de memória do i:", &i)

	/* a := &i
	fmt.Println("A variável a recebe o end de mem da variável i:", a)
	fmt.Println("A variável *a mostra o end de memória do a, e não o valor:", *a)
	fmt.Println("Variável a tem seu próprio endereço &a:", &a) */

	zeroValue(i)

	zeroPointer(&i)
	fmt.Println("zeroPointer:", i)
	fmt.Println("ponteiro:", &i)

}
