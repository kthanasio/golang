package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request Finalizada")
	select {
	case <-time.After(time.Second * 5):
		log.Println("Request Processada com Sucesso!")
		w.Write([]byte("Processamento Concluído com Sucesso!"))
	case <-ctx.Done():
		log.Println("Request Cancelada pelo Client")
	}
}
