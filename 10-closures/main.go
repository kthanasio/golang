package main

import (
	"fmt"
)

var (

)

func main() {
	n, t := sum(100, 50, 100)
	fmt.Printf("Soma: %d\nQuantidade: %d\n", n, t)

	total := func () int {
		val,_ := sum(100,50,100)
		return val * 2
	}()

	fmt.Printf("Dobro: %d\n", total)
}

func sum(numeros ...int) (int, int) {
	fmt.Println("Dentro do SUM")
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total, len(numeros)
}