package main

import (
	"fmt"
	"math"
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

func (r retangulo) area() float64 {
	return r.altura * r.largura

}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	c := circulo{raio: 10}
	r := retangulo{10, 15}

	areaCirculo := c.area()
	areaRetangulo := r.area()

	fmt.Println("Area Circulo: ", areaCirculo)
	fmt.Println("Area Retangulo: ", areaRetangulo)
}
