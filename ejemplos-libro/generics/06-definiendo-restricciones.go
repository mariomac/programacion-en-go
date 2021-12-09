package main

import (
	"fmt"
	"strings"
)

// a partir de interfaces comunes
func Mayusculas[T fmt.Stringer](elem T) string {
	return strings.ToUpper(elem.String())
}

type Persona struct {
	Nombre string
	Genero string
}

func (p Persona) String() string {
	return p.Nombre + " (" + p.Genero + ")"
}

// interfaces de restricciones
type Byte interface {
	int8 | uint8
}

func InvierteBits[T Byte](n T) T {
	var inv T
	for i := 0; i < 8; i++ {
		inv <<= 1
		inv |= n & 1
		n >>= 1
	}
	return inv
}

type Edad uint8

func main() {
	p := Persona{Nombre: "Lidia", Genero: "M"}
	fmt.Println(Mayusculas[Persona](p))
	fmt.Printf("%b\n", InvierteBits[uint8](0b01001100))
	fmt.Printf("%b\n", InvierteBits[int8](0b01001100))
	fmt.Printf("%b\n", InvierteBits[byte](0b01001100))

	// esto nos da un error
	//  Edad does not implement Byte (possibly missing ~ for uint8 in constraint Byte)
	// mostrar el nuevo tipo Byte "arreglado"
	// mostrar después todos los tipos del paquete constraints
	fmt.Printf("%b\n", InvierteBits[Edad](123))
}
