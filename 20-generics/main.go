package main

import "fmt"

// SEM GENERICS
// func SomaInt(m map[string]int) int {
// 	soma := 0
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }
// func SomaFloat(m map[string]float64) float64 {
// 	soma := 0.0
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// GENERICS - SEM CONSTRAINT
// func Soma[T int | float64] (m map[string]T) T {
// 	var soma T
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// GENERICS - COM CONSTRAINT
type Number interface {
	~int | ~float64
}

type MyNumber int

func Soma[T Number] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Comparar[T comparable] (a T, b T) bool {
	return a == b
}

func main() {
	salarios := map[string]int {"Kleber": 1000, "Viviane": 2000, "Gabriela": 3000}
	salarios2 := map[string]float64 {"Kleber": 10.0, "Viviane": 20.0, "Gabriela": 30.0}
	salarios3 := map[string]MyNumber {"Kleber": 1000, "Viviane": 2000, "Gabriela": 3000}
	// fmt.Printf("Salários = %v\n", SomaInt(salarios))
	// fmt.Printf("Salários2 = %v\n", SomaFloat(salarios2))
	fmt.Printf("Salários = %v\n", Soma(salarios))
	fmt.Printf("Salários2 = %v\n", Soma(salarios2))
	fmt.Printf("Salários3 = %v\n", Soma(salarios3))

	fmt.Println(Comparar[int](10, 10.0))
}