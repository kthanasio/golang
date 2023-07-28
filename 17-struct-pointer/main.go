package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c *Cliente) Andar() {
	c.Nome = "Kleber de Menezes"
	fmt.Printf("O cliente %v andou\n", c.Nome)
}

func main() {
	kleber := Cliente{
		Nome: "Kleber",
	}	
	kleber.Andar()
	fmt.Printf("Cliente: %v\n", kleber.Nome)
}