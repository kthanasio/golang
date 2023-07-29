package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		res, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var endereco ViaCEP
		err = json.Unmarshal(res, &endereco)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(endereco.Cep + ".txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = f.WriteString(fmt.Sprintf("CEP: %v;Localidade: %v;Logradouro: %v", endereco.Cep, endereco.Localidade, endereco.Logradouro))
		if err != nil {
			panic(err)
		}
	}

}
