package main

import (
	"fmt"

	"github.com/kthanasio/golang/37-APIS/configs"
	"github.com/kthanasio/golang/37-APIS/internal/entity"
	p "github.com/kthanasio/golang/37-APIS/pkg/entity"
)

func main() {
	_, err := configs.LoadConfig("./")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", config)

	kleber, err := entity.NewUser("Kleber", "klebermt@gmail.com", "pass@123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", kleber)

	valid := kleber.ValidatePassword("pass@123")
	fmt.Printf("Sucesso: %v\n", valid)

	validID, err := p.ParseID("31d0fd5d-85e0-4664-bb78-e1c704dea74f")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", validID)
	// -----

	camiseta, err := entity.NewProduct("Camiseta", 99.99)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", camiseta)
}
