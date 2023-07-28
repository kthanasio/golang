package main

import (
	"fmt"
)

var (

)

func main() {
	n, t := sum(100, 50, 100)
	fmt.Printf("Soma: %d\nQuantidade: %d\n", n, t)
}

func sum(numeros ...int) (int, int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total, len(numeros)
}