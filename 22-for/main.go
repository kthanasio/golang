package main

import "fmt"

func main() {
	
	for i:=0; i<3; i++ {
		fmt.Printf("Indice %v\n", i)
	}

	numeros := []string{"Kleber", "Leticia", "Gabriela"} // slice
	for i,v := range numeros {
		fmt.Printf("Indice %v e valor %v\n", i,v)
	}

	i := 0
	for i < 3 {
		fmt.Printf("Indice: %v\n", i)
		i++
	}
}