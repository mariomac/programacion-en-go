package main

import (
	"constraints"
	"fmt"
)

// usar "any" nos daría error
// invalid map key type T (missing comparable constraint)
func HayDuplicados[T comparable](elems ...T) bool {
	recorrido := map[T]struct{}{}
	for _, e := range elems {
		if _, ok := recorrido[e]; ok {
			return true
		}
		recorrido[e] = struct{}{}
	}
	return false
}

// explicar demás elementos del paquete constraints
func Maximo[T constraints.Ordered](primero T, elems ...T) T {
	max := primero
	for _, e := range elems {
		if e > max {
			max = e
		}
	}
	return max
}

func main() {
	fmt.Println(HayDuplicados[int](1, 2, 3, 2, 4))
	fmt.Println(HayDuplicados[string]("hola", "que", "tal"))
	fmt.Println(Maximo[int](1, 2, 3, 2, 1, 1))
	fmt.Println(Maximo[float32](1.0, 1.1, 1.15, 0.7))
	fmt.Println(Maximo[string]("aaa", "abc", "cda"))
}
