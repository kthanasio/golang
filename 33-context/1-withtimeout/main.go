package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Printf("Execução cancelada!\n")
	case <-time.After(time.Second * 2):
		fmt.Printf("Execução Finalizada com Sucesso!\n")
	}
}
