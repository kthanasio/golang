package main

import "fmt"

func main() {
	var x interface{} = "Kleber"
	fmt.Println(x.(string))

	var n interface{} = "10"
	res, ok := n.(int)
	fmt.Printf("O valor de res é %v e o resultado de OK é %v\n",res, ok)
}