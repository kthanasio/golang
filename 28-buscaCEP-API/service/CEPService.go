package service

import (
	"28-buscaCEP-API/model"
	"encoding/json"
	"io"
	"net/http"
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
