package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println(ptr)

	val := 5
	var ptr *int = &val

	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(*ptr + 1)
}