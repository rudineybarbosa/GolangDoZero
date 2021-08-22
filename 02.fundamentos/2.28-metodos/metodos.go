package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) imprimirDados() {
	fmt.Printf("Nome: %s, Idade: %d \n", u.nome, u.idade)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario2 := usuario{nome: "Aline", idade: 35}
	usuario := usuario{nome: "Rudiney", idade: 38}

	usuario.imprimirDados()
	usuario2.imprimirDados()

	usuario2.fazerAniversario()
	usuario2.imprimirDados()
}
