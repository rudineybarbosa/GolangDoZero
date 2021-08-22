package main

import "fmt"

var n int

func init() {
	fmt.Println("1 - Primeira função a ser executada. Bom para iniciar valores")
	n = 10
}

func main() {
	fmt.Println("2 - Função main")
	fmt.Println(n)
}
