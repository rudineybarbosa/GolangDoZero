package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10

	//passagem de valor por COPIA
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	//PONTEIROS é um endereço de memória
	var variavel3 int
	var variavelPonteiro4 *int

	variavel3 = 100
	//ERRO NA LIHA ABAIXO
	//Um ponteiro não recebe um valor, ele deve receber um endereço de memória
	//variavelPonteiro4 = variavel3

	//O ponteiro recebendo um endereço de memória
	variavelPonteiro4 = &variavel3
	fmt.Println(variavel3, variavelPonteiro4) //EXEMPLO de resultado: 0xc000018108

	//Para obter o valor que e encontra no endereço de memória do ponteiro, utiliza-se asteristico antes da variável
	fmt.Println(variavel3, *variavelPonteiro4)

	//Alterando o valor que se encontra no endereço de memória do ponteiro
	variavel3 = 150
	fmt.Println(variavel3, *variavelPonteiro4)

}
