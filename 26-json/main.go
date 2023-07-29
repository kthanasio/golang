package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"id"`
	Saldo  int `json:"saldo"`
	Banco  int `json:"-"`
}

func main() {

	conta := Conta{
		Numero: 1,
		Saldo:  1000,
		Banco:  1,
	}
	// res, err := json.Marshal(conta)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(res))
	err := json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"id":1,"saldo":1100,"banco":1}`)
	var contaKleber Conta
	err = json.Unmarshal(jsonPuro, &contaKleber)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saldo: %v\n", contaKleber.Saldo)
}
