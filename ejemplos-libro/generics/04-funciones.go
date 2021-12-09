package main

import "fmt"

func Parte[T any](elementos []T) ([]T, []T) {
	mitad := len(elementos) / 2
	return elementos[0:mitad], elementos[mitad:]
}

func main() {
	fmt.Println(Parte([]int{1, 2, 3}))
}
