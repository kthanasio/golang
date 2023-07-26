package main

import (
	"fmt"
)

var (
	
)

func main() {
	s := []int{100,200,300,400,500,600,700}
	fmt.Printf("len=%d cap=%d val=%v \n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d val=%v \n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d val=%v \n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len=%d cap=%d val=%v \n", len(s[2:]), cap(s[2:]), s[2:])
	s = append(s, 800)
	fmt.Printf("\nlen=%d cap=%d val=%v \n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d val=%v \n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d val=%v \n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len=%d cap=%d val=%v \n", len(s[2:]), cap(s[2:]), s[2:])

	for i,v := range s {
		fmt.Printf("len=%d ind=%d val=%d \n", len(s), i, v)
	}
}