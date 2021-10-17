package main

import (
	"fmt"
	"time"
)

func main() {

	//Cria um canal
	channel := make(chan string)

	go escrever("Rotina 01", channel)

	//SINTAXE PARA O SISTEMA PARAR E FICAR AGUARDANDO O CANAL RECEBER UMA INFORMAÇÃO DO TIPO INDICADO NA SUA DECLARAÇÃO
	//O FLUXO SEGUE APENAS APÓS UMA INFORMAÇÃO FOR ENVIADA PARA O CANAL E O TRECHO ABAIXO RECEBER O DADO
	// texto := <-channel
	for texto := range channel {
		fmt.Println(texto)

	}

}

func escrever(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- texto

		time.Sleep(time.Second)
	}

	close(channel)
}
