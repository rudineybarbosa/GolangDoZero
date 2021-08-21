package main

import "fmt"

func main() {
	var nome string = "str1"
	fmt.Println(nome)

	//tipo implicito
	str2 := "str2"
	fmt.Println(str2)

	var (
		str3 string = "str3"
		str4 string = "str4"
	)
	fmt.Println(str3, str4)

	str5, numero6 := "str5", 6
	fmt.Println(str5, numero6)

	//MESMA REGRA ACIMA PARA INICIAR CONSTANTES
	//POREM UTILIZA-SE A PALAVRA const

	const str7 = "str7"
	fmt.Println(str7)
}
