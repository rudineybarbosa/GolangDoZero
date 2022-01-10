package main

import (
	"fmt"
)

func main() {

	//CANAL DE STRING COM BUFFER DE TAMANHO 2
	canal := make(chan string, 2)

	//SE O CANAL RECEBER MAIS QUE 2 STRINGS, TEREMOS DEADLOCK PORQUE O PROGRAMA NUNCA RECEBERÁ UMA TERCEIRA STRING
	canal <- "Ola 1"
	canal <- "Ola 2"

	mensage := <-canal
	mensage2 := <-canal

	//SE O CANAL TENTAR ENVIAR UMA TERCEIRA STRING, TAMBÉM TEREMOS DEADLOCK NO PROGRAMA
	//mensage3 := <-canal

	fmt.Println(mensage)
	fmt.Println(mensage2)

}
