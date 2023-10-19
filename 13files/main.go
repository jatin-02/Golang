package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "To be written in the file"
	file, err := os.Create("./file1.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println(length)
	defer file.Close()

	readFile("./file1.txt")
}

func readFile(filename string)  {
	dataBytes, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println(dataBytes)
	fmt.Println(string(dataBytes))
}

func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}