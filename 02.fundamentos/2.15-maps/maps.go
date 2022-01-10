package main

import (
	"fmt"
)

type endereco struct {
	rua    string
	numero uint8
	bairro string
	cidade string
	estado string
}

func main() {

	//DECLARAÇÃO DE MAPs:
	//map[tipo da chave]tido do valor{}
	usuario := map[string]string{
		"nome":      "Rudi",
		"sobrenome": "Barbosa",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	catalogo := map[string]endereco{
		"luzinete": {"rua teles de menezes", 231, "jardim silvina", "SBC", "São Paulo"},
		"rudiney":  endereco{"rua dois de julho", 251, "santo amaro", "Recife", ""},
	}
	fmt.Println(catalogo)
	fmt.Println(catalogo["rudiney"])
	fmt.Println(catalogo["rudiney"].cidade)
}
