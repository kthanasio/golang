package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kthanasio/golang/cep/model"
)

// CEPService will call an external API to ge data
func CEPService(cep string) (*model.ViaCEP, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var endereco model.ViaCEP
	err = json.Unmarshal(res, &endereco)
	if err != nil {
		return nil, err
	}
	return &endereco, nil
}
