package main

import "fmt"

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
func Soma[T int | float64] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	salarios := map[string]int {"Kleber": 1000, "Viviane": 2000, "Gabriela": 3000}
	salarios2 := map[string]float64 {"Kleber": 10.0, "Viviane": 20.0, "Gabriela": 30.0}
	// fmt.Printf("Salários = %v\n", SomaInt(salarios))
	// fmt.Printf("Salários2 = %v\n", SomaFloat(salarios2))
	fmt.Printf("Salários = %v\n", Soma(salarios))
	fmt.Printf("Salários2 = %v\n", Soma(salarios2))
}