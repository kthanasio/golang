package main

import (
	"21-pacotes-modules/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New()
	fmt.Printf("Novo id: %v\n", uuid)
	valor1 := 10.0
	valor2 := 20.0
	fmt.Printf("Soma: %v\n", matematica.Somar(valor1, valor2))
	fmt.Printf("Subtrai: %v\n", matematica.Subtrair(valor1, valor2))
	fmt.Printf("Multiplica: %v\n", matematica.Multiplicar(valor1, valor2))
	fmt.Printf("Divide: %v\n", matematica.Dividir(valor1, valor2))

}