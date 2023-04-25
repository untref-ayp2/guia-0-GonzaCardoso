package main

import (
	"fmt"
)

func main() {
	var funcion []float32 = []float32{-2.0, -4.00, -0.52, 0.00, 0.03, 1.50}
	var i int = 0
	if funcion[0] != 0.0 {
		fmt.Print(funcion[0])
	}

	for i < len(funcion)-1 {
		if funcion[i+1] < 0.0 {
			fmt.Print(" ", funcion[i+1], "x^", i+1)
			i++
		} else if funcion[i+1] == 0.0 {
			i++
		} else {
			fmt.Print(" +", funcion[i+1], "x^", i+1)
			i++
		}
	}
}
