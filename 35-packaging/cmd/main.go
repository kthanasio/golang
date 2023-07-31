package main

import (
	"fmt"

	"github.com/kthanasio/golang/35-packaging/math"
)

func main() {
	m := math.Math{A: 10, B: 20}
	r, err := m.Add()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resultado: %v\n", r)
}
