package main

import "fmt"

func soma(a , b *int) int {
	*a = 50
	*b = 50
	fmt.Printf("a endereco: %p - a valor: %d\n", &a, *a)
	fmt.Printf("b endereco: %p - b valor: %d\n", &b, *b)
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	fmt.Printf("minhaVar1 endereco: %p - minharVar1 valor: %d\n", &minhaVar1, minhaVar1)
	fmt.Printf("minhaVar2 endereco: %p - minharVar2 valor: %d\n", &minhaVar2, minhaVar2)

	fmt.Printf("Total %d\n", soma(&minhaVar1, &minhaVar2))

	fmt.Printf("minhaVar1 endereco: %p - minharVar1 valor: %d\n", &minhaVar1, minhaVar1)
	fmt.Printf("minhaVar2 endereco: %p - minharVar2 valor: %d\n", &minhaVar2, minhaVar2)
}