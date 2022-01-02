package main

import (
	"fmt"
	"strings"
)

func JuntaStrings[T fmt.Stringer](elems ...T) string {
	str := strings.Builder{}
	for _, e := range elems {
		str.WriteString(e.String())
	}
	return str.String()
}

func JuntaStrings2(elems ...fmt.Stringer) string {
	str := strings.Builder{}
	for _, e := range elems {
		str.WriteString(e.String())
	}
	return str.String()
}

type Persona struct {
	Nombre string
	Genero string
}

func (p Persona) String() string {
	return p.Nombre + " (" + p.Genero + ")"
}

// interfaces de restricciones
type OchoBits interface {
	~int8 | ~uint8
}

func InvierteBits[T OchoBits](n T) T {
	var inv T
	for i := 0; i < 8; i++ {
		inv <<= 1
		inv |= n & 1
		n >>= 1
	}
	return inv
}

type Edad uint8

func (t Edad) String() string {
	return fmt.Sprint(uint8(t))
}

func main() {
	p := Persona{Nombre: "Lidia", Genero: "M"}
	fmt.Println(JuntaStrings(p, p))
	fmt.Println(JuntaStrings2(p, Edad(33)))
	fmt.Printf("%08b\n", InvierteBits[uint8](0b01001000))
	fmt.Printf("%08b\n", InvierteBits(int8(0b01100100)))
	fmt.Printf("%b\n", InvierteBits(Edad(33)))

}
