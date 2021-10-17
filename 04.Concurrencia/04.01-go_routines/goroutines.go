package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("Rotina 01")
	escrever("Rotina 02")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
