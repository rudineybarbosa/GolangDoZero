package main

import (
	"fmt"
	"math"
	"reflect"
)

type forma interface {
	area() float64
}

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

//Method 'area' deined to retangulo structure
func (r retangulo) area() float64 {
	return r.altura * r.largura

}

//Method 'circulo' deined to retangulo structure
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func calcularArea(f forma) {

	if reflect.TypeOf(f).Name() == "circulo" {
		fmt.Println("Area Círculo: ", f.area())

	} else {
		fmt.Println("Area Retângulo: ", f.area())

	}

}

func main() {
	c := circulo{raio: 10}
	r := retangulo{10, 15}

	areaCirculo := c.area()
	areaRetangulo := r.area()
	fmt.Println("Area Circulo: ", areaCirculo)
	fmt.Println("Area Retangulo: ", areaRetangulo)

	//UTILIZANDO FUNÇÃO QUE RECEBE TIPO GENÉRICO
	calcularArea(c)
	calcularArea(r)

}
