package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	apiMaxTime = 200
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request Iniciada - Timeout[%+v]\n", time.Millisecond*time.Duration(apiMaxTime))
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(apiMaxTime))
	response, err := Executa(&ctx)
	if err != nil {
		log.Fatalf("Timeout na chamada da API")
	}
	defer cancel()

	fmt.Printf("Resultado: %v\n", response)
	defer log.Println("Request Finalizada")
}

func Executa(ctx *context.Context) (int, error) {
	log.Println("Vai Chamar a API")
	req, err := http.NewRequestWithContext(*ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return 1, nil
}
