package main

import "fmt"

func main() {
	defer fmt.Println("No 1")
	defer fmt.Println("No 2")
	defer fmt.Println("No 3")
	fmt.Println("No 4")
}