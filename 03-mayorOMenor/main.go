package main

import (
	"fmt"
)

func main() {
	var lista []int = []int{-700, -4, -4, -23, -59, -5, -6}
	fmt.Print("mayor = ", elMayor(lista), " ")
	fmt.Print("menor = ", elMenor(lista))

}

func elMayor(z []int) int {
	var mayor int = z[0]
	for i := 1; i < len(z)-1; i++ {
		if z[i] > mayor {
			mayor = z[i]
		}
	}
	return mayor
}

func elMenor(z []int) int {
	var menor int = z[0]
	for i := 1; i < len(z)-1; i++ {
		if z[i] < menor {
			menor = z[i]
		}
	}
	return menor
}
