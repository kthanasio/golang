package main

import (
	"fmt"
)

var (

)

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	Endereco // composição
	Documento Documento // propriedade
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("Cliente %s desativado\n", c.Nome)
}
type Endereco struct {
	Logradouro string
	Numero string
	Cidade string
	Estado string
}

type Documento struct {
	Tipo string
	Numero string
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}
func main() {
	kleber := Cliente{
		Nome: "Kleber de Menezes Thanasio",
		Idade: 39,
		Ativo: true,
		Documento: Documento{ Tipo: "CPF", Numero: "123.123.123-00"},
	}

	kleber.Logradouro = "Rua Teste"
	kleber.Numero = "87"
	kleber.Cidade = "São Paulo"
	kleber.Estado = "SP"

	// kleber.Desativar()
	Desativacao(kleber)

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\nEndereco: %s\nDocumento: %s - %s\n",
				kleber.Nome,
				kleber.Idade,
				kleber.Ativo,
				kleber.Endereco.Logradouro + ", " + kleber.Endereco.Numero,
				kleber.Documento.Tipo,
				kleber.Documento.Numero)
	
}