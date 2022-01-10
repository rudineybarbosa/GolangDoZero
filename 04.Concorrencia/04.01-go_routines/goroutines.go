package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("Rotina 01") //inicia a execução da rotina 1, e uma outra thread segue para próxima linha
	escrever("Rotina 02")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
