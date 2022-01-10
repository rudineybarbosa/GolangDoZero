package main

import "fmt"

func main() {

	var array1 [5]int
	array1[1] = 1
	fmt.Println("Array1: ", array1)

	var array2 [5]string
	fmt.Println("Array2: ", array2)

	array3 := [3]string{"Pos1", "Pos2", "Pos3"}
	fmt.Println("Array3: ", array3)

	//ARRAY DINÂMICO
	array4 := [...]int{1, 2, 3, 4}
	fmt.Println(array4)

	fmt.Println("SLICE	###########################################")
	//Não é um array. Ele é um ponteiro para um array interno, que é criado em tempo de execução
	slice1 := []int{1, 2, 4, 5, 6, 7}
	fmt.Println(slice1)
	slice1 = append(slice1, 3)
	fmt.Println(slice1)

	//slice pode ser criado como uma parte de um array
	//E podemos comprovar que é um ponteiro fazendo uma alteração no array original
	slice2 := array3[0:2]
	fmt.Println("slice 2: ", slice2)
	array3[1] = "Pos2 Alterada"
	fmt.Println("slice 2 pos alteração: ", slice2)
	fmt.Println("Array3: ", array3)

	//SLICE é um array interno criado em tempo de execução. MAs pode ser criado explicitamente
	//Primeiro argumento é o tipo dos elementos armazenados;
	//O segundo argumento é o tamanho do slice retornado DISPONIVEL para uso;
	//O terceiro argumento é a capacidade total do slice.
	//Caso a capacidade seja ultrapassada, o sistema cria outro slice com o dobro do espaço que foi ultrapassado
	slice3 := make([]int, 3, 5)
	fmt.Println("slice3: ", slice3)
	fmt.Println("length slice3: ", len(slice3))
	fmt.Println("capacidade slice3: ", cap(slice3))
	slice3 = append(slice3, 1, 2, 3, 4)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
