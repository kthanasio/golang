package main

import "fmt"

func main() {
	
	idade := 18
	sexo := "Male"

	if sexo == "Male" && (idade >= 18 && idade <= 60) {
		fmt.Printf("Alistar no exercito - Idade [%v] - Sexo [%v]\n", idade, sexo)
	} else {
		fmt.Printf("Dispensado - Idade [%v] - Sexo [%v]\n", idade, sexo)
	}

	if idade < 18 || idade > 60 {
		fmt.Printf("Dispensado - Idade [%v] - Sexo [%v]\n", idade, sexo)
	}

	switch sexo {
	case "Female":
		fmt.Printf("Dispensado - Sexo [%v]\n", sexo)
	case "Male":
		fmt.Printf("Alistar no exercito - Sexo [%v]\n", sexo)
	default:
		fmt.Printf("Outro\n")
	}
}