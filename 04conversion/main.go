package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome...")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please give us a rating")

	input, _ := reader.ReadString('\n')
	updatedRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The updated rating is, ", updatedRating+1)
	}
}