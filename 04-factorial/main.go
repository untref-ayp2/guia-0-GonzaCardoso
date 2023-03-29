package main

import "fmt"

func main() {
	fmt.Print(factor(3))
}
func factor(z int) int {
	resultado := 1
	for i := 2; i <= z; i++ {
		resultado = i * resultado
	}
	return resultado
}
