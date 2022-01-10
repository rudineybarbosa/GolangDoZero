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
	//O FLUXO SEGUIRÁ APENAS APÓS UMA INFORMAÇÃO FOR ENVIADA PARA O CANAL E O TRECHO ABAIXO RECEBER O DADO
	//texto := <-channel

	//OUTRA FORMA DE FICAR LENDO DADOS DO CANAL, É COM UM 'for', PORQUE O CANAL TERÁ UM CONJUNTO DE ENTRADAS E ASSIM PODEMOS FAZER ITERAÇÃO
	for texto := range channel {
		fmt.Println(texto)

	}

}

//variável channel, do tipo chan, para trafegar string
func escrever(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- texto

		time.Sleep(time.Second)
	}

	//FECHANDO O CANAL
	close(channel)
}
