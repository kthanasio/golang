package main

import "fmt"

func showType(t interface{}) {
	fmt.Printf("O Tipo da variável t é %T e o valor é %v\n", t, t)
}

func main() {
	var x interface{} = 10
	showType(x)
	x = "Testing"
	showType(x)
	var y interface{} = "Hello, World"
	showType(y)
	y = true
	showType(y)
}