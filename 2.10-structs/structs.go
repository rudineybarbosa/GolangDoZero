package main

import "fmt"

//struct: conceito mais próximo de Classe
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	var u usuario
	u.nome = "Rudiney"
	u.idade = 38

	//não dá erro de execução ao criar um
	//struct com a sintaxe acima, faltando
	//dados internos, como endereco
	fmt.Println(u)

	address := endereco{"rua teles de meneses", 231}

	//Já com esta sintaxe de inferência, devemos
	//informar todos os atributos do struct, caso
	// contrário da erro
	usuario2 := usuario{"Rudiney", 38, address}
	fmt.Println(usuario2)

}
