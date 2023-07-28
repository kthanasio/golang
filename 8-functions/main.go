package main

import (
	"errors"
	"fmt"
)

var (

)

func main() {
	// fmt.Println(sum(1, 20))
	// fmt.Println(sum2(1, 20))
	v, err := sum3(1, 22)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a, b int) int {
	return a + b
}

func sum3(a, b int) (int, error) {
	if (a+b) > 10 {
		return a + b, errors.New("Soma > 10")	
	}
	return a + b, nil
}