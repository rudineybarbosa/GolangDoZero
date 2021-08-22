package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
		NUMBERS
	*/

	//INTEIROS
	//int8, int16, int32, int64
	//Aceitam valores negativos

	//tipo impĺicito para num0 e num1
	num0 := 2134214321
	var num1 = 1000000000
	fmt.Println(num0)
	fmt.Println(num1)

	//tipo explicito
	var num2 int32 = 1000000000
	fmt.Println(num2)

	//o tipo to int será inferido de acordo com
	//a arquitetura do SO em execução.
	//Neste caso, será int64
	var num3 int = 1000000000
	fmt.Println(num3)

	//FLOATS
	//Abaixo nao compila pois para floats
	//precisamos informar o tipo do float
	//float32, float64
	//var num4 float = 123.30
	//fmt.Println(num4)

	//tipo implicito funciona com float: num5 e num6
	num6 := 1234.43
	var num5 = 123.12
	fmt.Println(num5)
	fmt.Println(num6)

	var num7 float32 = 1234.123
	var num8 float64 = 2345.23
	fmt.Println(num7)
	fmt.Println(num8)

	//uint8, uint16, uint32, uint64
	//Somente valores positivos (u: unsigned)
	//O codigo abaixo não compila
	//var num9 uint8 = -123

	var (
		num10 int64   = 1234567890
		num11 float64 = 1234.123
	)
	fmt.Println(num10, num11)

	/*
		STRINGS
	*/
	var str1 string = "str1"
	fmt.Println(str1)

	//tipo implicito
	str2 := "str2"
	fmt.Println(str2)

	var (
		str3 string = "str3"
		str4 string = "str4"
	)
	fmt.Println(str3, str4)

	str5, str6 := "str5", "str6"
	fmt.Println(str5, str6)

	/*
		OBTER VALOR TABELA ASC
	*/
	var (
		a = 'a'
		b = 'b'
		c = 'c'
	)
	fmt.Print("Valores tabela ASCII para 'a', 'b' e 'c': ")
	fmt.Println(a, b, c)

	/*
	 BOOLEANOS
	*/
	var bool1 bool = true
	fmt.Println(bool1)

	/*
	 Tipo ERROR
	 Valor default é nil
	*/
	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Erros customizado")
	fmt.Println(erro2)
}
