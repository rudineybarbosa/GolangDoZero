package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//GO permite funcao com mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("Funcao f")
		fmt.Println(txt)

		return "fim funcao f"
	}

	resultado := f("texto da funcao f")
	println(resultado)

	resultadosSoma, resultadoSubtracao := calculosMatematicos(10, 20)
	println(resultadosSoma, resultadoSubtracao)

	//CASO NAO PRECISE DE UM DOS RETORNOS, ATRIBUA A UNDERLINE PARA IGNORAR
	resultadosSoma2, _ := calculosMatematicos(10, 20)
	_, resultadoSubtracao2 := calculosMatematicos(10, 20)
	println(resultadosSoma2)
	println(resultadoSubtracao2)

}
