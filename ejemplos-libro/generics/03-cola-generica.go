package main

import "fmt"

type Persona struct {
	Nombre string
	Genero string
}

type Cola[T any] struct {
	elementos []T
}

func (c *Cola[T]) Introduce(elem T) {
	c.elementos = append(c.elementos, elem)
}

func (c *Cola[T]) Extrae() T {
	if len(c.elementos) == 0 {
		var valorCero T
		return valorCero
	}
	ret := c.elementos[0]
	c.elementos = c.elementos[1:]
	return ret
}

func main() {
	tanda := Cola[Persona]{}
	tanda.Introduce(Persona{Nombre: "Mari Carmen", Genero: "M"})
	tanda.Introduce(Persona{Nombre: "Juan", Genero: "H"})

	// esto daría error de compilación
	// cannot use "Fernando" (untyped string constant) as Persona value in argument to tanda.Introduce
	//tanda.Introduce("Fernando")

	fmt.Println("despachando a", tanda.Extrae().Nombre)
	fmt.Println("despachando a", tanda.Extrae().Nombre)

	codigos := Cola[int]{}
	codigos.Introduce(6)
	codigos.Introduce(7)
	codigos.Introduce(8)

	var siguiente int = codigos.Extrae()

	_ = siguiente

}
