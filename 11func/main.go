package main

import "fmt"

func main() {
	addResult := adder(3, 5)
	fmt.Println(addResult)

	proResult := proAdder(1, 2, 3, 4, 5)
	fmt.Println(proResult)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range(values) {
		total += value
	}

	return total
}