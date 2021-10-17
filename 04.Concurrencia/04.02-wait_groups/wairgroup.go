package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var group sync.WaitGroup

	group.Add(2) //indica total de grupos

	go func() {
		escrever("Rotina 01")

		group.Done() //reduz 1 do total de grupos
	}()

	go func() {
		escrever("Rotina 02")

		group.Done() //reduz 1 do total de grupos
	}()

	group.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
