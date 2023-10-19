package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev/doc/"

func main() {
	response, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Response type is %T\n", response)

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	checkNilError(err)

	fmt.Println(string(dataBytes))
}

func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}