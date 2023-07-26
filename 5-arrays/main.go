package main

import (
	"fmt"
)

var (
	// array tem tamanho fixo
	minhaLista [3]int
)

func main() {
	minhaLista[0] = 100
	minhaLista[1] = 200
	minhaLista[2] = 300
	fmt.Printf( "%v\n", minhaLista[0])
	fmt.Printf( "%v\n", minhaLista[len(minhaLista)-1])

	for i, v := range minhaLista {
		fmt.Printf("Indice [%d] valor [%d]\n", i, v)
	}
}