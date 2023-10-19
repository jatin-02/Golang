package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome...")
	
	fmt.Println("Please give us a rating")
	reader := bufio.NewReader(os.Stdin)

	// comma ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating, ", input)
	fmt.Printf("Input is of type %T", input)
}