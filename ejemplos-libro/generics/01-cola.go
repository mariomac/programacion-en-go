package main

import "fmt"

type Persona struct {
	Nombre string
	Genero string
}

type Cola struct {
	personas []Persona
}

func (c *Cola) Introduce(elem Persona) {
	c.personas = append(c.personas, elem)
}

func (c *Cola) Extrae() Persona {
	if len(c.personas) == 0 {
		return Persona{}
	}
	ret := c.personas[0]
	c.personas = c.personas[1:]
	return ret
}

func main() {
	tanda := Cola{}
	tanda.Introduce(Persona{Nombre: "Juan", Genero: "H"})
	tanda.Introduce(Persona{Nombre: "Maria del Carmen", Genero: "M"})
	tanda.Introduce(Persona{Nombre: "Juan Carlos", Genero: "H"})

	fmt.Println("despachando a", tanda.Extrae().Nombre)
	fmt.Println("despachando a", tanda.Extrae().Nombre)

	tanda.Introduce(Persona{Nombre: "Eduardo", Genero: "H"})

	fmt.Println("despachando a", tanda.Extrae().Nombre)
	fmt.Println("despachando a", tanda.Extrae().Nombre)
}
