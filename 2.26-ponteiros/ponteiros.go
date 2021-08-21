package main

import "fmt"

func inverterNumero(numero int) int {
	numero = numero * -1
	return numero
}

func inverterNumeroPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 10
	numeroInvertido := inverterNumero(numero)
	fmt.Println(numero)
	fmt.Println(numeroInvertido)

	fmt.Println("Agora utilizando ponteiro, a referência original será alterada")
	num := 50
	inverterNumeroPonteiro(&num)
	fmt.Println(num)

}
