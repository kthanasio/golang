package main

import (
	"fmt"
)

var (

)

func main() {
	salarios := map[string]float64 {"Kleber": 1000.0, "Viviane": 2000.0, "Marcel": 3000.0}
	// fmt.Println(salarios["Kleber"])
	// fmt.Println(salarios["Viviane"])
	// fmt.Println(salarios["Marcel"])
	salarios["Leticia"] = 4000.0
	salarios["Gabriela"] = 5000.0
	// fmt.Println(salarios["Leticia"])
	// fmt.Println(salarios["Gabriela"])

	for nome, salario := range salarios {
		fmt.Printf("Nome [%v] e salário [%f]\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("Salário [%f]\n", salario)
	}
	// sal := make(map[string]int)
	// fmt.Println(sal)

}