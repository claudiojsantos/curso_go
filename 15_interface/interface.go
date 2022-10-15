package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área do objeto é %0.2f\n", f.area())
}

type retangulo struct {
	base   float64
	altura float64
}

func (r retangulo) area() float64 {
	return r.base * r.altura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{20, 30}
	escreverArea(r)

	c := circulo{30}
	escreverArea(c)
}
