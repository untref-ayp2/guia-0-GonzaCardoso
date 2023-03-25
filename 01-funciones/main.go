package main

import (
	"fmt"
)

func main() {
	var funcion []float64 = []float64{3.43, -4.00, -0.52, 0.00, 0.03, 1.50}
	var i int = 0
	var h int = 1
	fmt.Print(funcion[0])
	for i < len(funcion)-1 {
		if funcion[h] < 0.0 {
			fmt.Print(" ", funcion[h], "*x^", h)
			i++
			h++
		} else if funcion[h] == 0.0 {
			i++
			h++
		} else {
			fmt.Print(" +", funcion[h], "*x^", h)
			i++
			h++

		}
	}
}
