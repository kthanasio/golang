package matematica

import "fmt"

func Somar[T int | float64] (a T, b T) T {
	logging("Chamou Somar")
	return a + b
}

func Subtrair[T int | float64] (a T, b T) T {
	logging("Chamou Subtrair")
	return a - b
}

func Multiplicar[T int | float64] (a T, b T) T {
	logging("Chamou Multiplicar")
	return a * b
}

func Dividir[T int | float64] (a T, b T) T {
	if (b == 0) { return 0 }
	return a / b
}

func logging(v string) {
	fmt.Printf("Logging: %v\n", v)
}