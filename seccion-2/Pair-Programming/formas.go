package main

import "fmt"

func main() {
	miTriangulo := triangulo{base: 2.7, altura: 5.6}
	miRectangulo := rectangulo{base: 10.6, altura: 9.8}
	miCirculo := circulo{radio: 1.2}
	miTrapecio := trapecio{base1: 2.2, base2: 4.5, altura: 8.1}
	miPoligono := poligono{apotema: 10.8, perimetro: 15.3}

	objetos := make(map[string]Shapes)

	objetos["Triangulo"] = miTriangulo
	objetos["Rectangulo"] = miRectangulo
	objetos["Circulo"] = miCirculo
	objetos["Trapecio"] = miTrapecio
	objetos["Poligono"] = miPoligono

	for _, o := range objetos {
		imprimirArea(o)
	}
}

func imprimirArea(s Shapes) {
	if s.area() > 100 {
		panic("Error, el area es demasiado grande")
	} else {
		fmt.Println(s.area())
	}
	defer func() {
		cadena := recover()
		if cadena != nil {
			fmt.Println(cadena)
		}
	}()
}

type triangulo struct {
	base   float64
	altura float64
}

func (t triangulo) area() float64 {
	return (t.base * t.altura) / 2
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

type circulo struct {
	radio float64
}

func (ci circulo) area() float64 {
	return 3.14 * (ci.radio * ci.radio)
}

type trapecio struct {
	base1  float64
	base2  float64
	altura float64
}

func (t trapecio) area() float64 {
	return ((t.base1 + t.base2) / 2) * t.altura
}

type poligono struct {
	lado      int
	apotema   float64
	perimetro float64
}

func (p poligono) area() float64 {
	return (p.apotema * p.perimetro) / 2
}

type Shapes interface {
	area() float64
}
