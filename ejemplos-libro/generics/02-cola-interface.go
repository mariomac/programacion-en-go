package main

import "fmt"

type Persona struct {
	Nombre string
	Genero string
}

type Cola struct {
	elementos []interface{}
}

func (c *Cola) Introduce(elem interface{}) {
	c.elementos = append(c.elementos, elem)
}

func (c *Cola) Extrae() interface{} {
	if len(c.elementos) == 0 {
		return ""
	}
	ret := c.elementos[0]
	c.elementos = c.elementos[1:]
	return ret
}

func main() {
	tanda := Cola{}
	tanda.Introduce(Persona{Nombre: "Juan", Genero: "H"})
	tanda.Introduce("Fernando")

	fmt.Println("despachando a", tanda.Extrae().(Persona).Nombre)
	fmt.Println("despachando a", tanda.Extrae().(Persona).Nombre)

}
