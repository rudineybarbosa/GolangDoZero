package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

//FORMA DE DIZER QUE UM ESTUDANTE É UMA PESSOA
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("Herança")

	p1 := pessoa{"Rudiney", "Barbosa", 38, 170}
	fmt.Println(p1)

	//PARA CRIAR UM ESTUDANTE, PRECISAMOS DE UMA PESSOA
	e1 := estudante{p1, "Computação", "UFPE"}
	fmt.Println(e1)

	/*PARA ACESSAR UMA VARIÁVEL DE UMA
	ESTRUTURA INTERNA, NÃO PRECISAMOS NAVEGAR
	ATRAVÉS DA MESMA. PODEMOS CHAMAR DIRETO O NOME DA PROPRIDADE
	*/
	fmt.Println(e1.nome)
}
