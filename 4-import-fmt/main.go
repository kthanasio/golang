package main

import (
	"fmt"
)

type ID int

var (
	a int = 123
	numero ID = 999
)
func main() {
	fmt.Printf("O tipo da variável é: %T", numero)
}