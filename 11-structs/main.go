package main

import (
	"fmt"
)

var (

)

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	kleber := Cliente{
		Nome: "Kleber de Menezes Thanasio",
		Idade: 39,
		Ativo: true,
	}
	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", kleber.Nome, kleber.Idade, kleber.Ativo)
	kleber.Ativo = false
	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", kleber.Nome, kleber.Idade, kleber.Ativo)
}