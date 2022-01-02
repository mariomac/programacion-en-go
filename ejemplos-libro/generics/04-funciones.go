package main

import "fmt"

func Partir[T any](elementos []T) ([]T, []T) {
	mitad := len(elementos) / 2
	return elementos[0:mitad], elementos[mitad:]
}

func main() {
	izda, dcha := Partir[int]([]int{1, 2, 3, 4})
	fmt.Println("parte izquierda", izda)
	fmt.Println("parte derecha", dcha)
	// Es mejor así, porque hay inferencia de tipos
	fmt.Println(Partir([]int{1, 2, 3, 4}))
	fmt.Println(Partir([]string{"a", "b", "c", "d"}))
}
