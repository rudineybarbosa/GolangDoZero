package main

func main() {
	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 4

	println(soma, subtracao, divisao,
		restoDivisao, multiplicacao)

	//GO é fortemente tipado e não permite operações
	//com tipos diferentes
	/*
		var num1 int16 = 10
		var num2 int32 = 20
		soma2 := num1 + num2
		println(soma2)
	*/
	//FIM DOS ARITMETICOS

	//ATRIBUIÇÃO
	var variavel1 string = "teste"
	variavel2 := 10
	println(variavel1, variavel2)
	//FIM ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	println(1 > 2)
	println(1 >= 2)
	println(1 == 2)
	println(1 == 1)
	println(1 <= 2)
	println(1 < 2)
	println(1 > 2)
	println(1 != 2)

	//OPERADORES LOGICOS
	println("-------------------")
	verdadeiro, falso := true, false
	println(verdadeiro && falso)
	println(verdadeiro || false)
	println(!verdadeiro)
	println(!falso)

	//OPERADORES UNARIOS
	numero := 10
	numero++
	println(numero)
	numero += 10
	println(numero)
	numero--
	println(numero)
	numero -= 5
	println(numero)
	numero *= 2
	println(numero)
	numero %= 3
	println(numero)

	//OPERADOR TERNARIO '?'
	//NÃO EXISTE EM GO
}
