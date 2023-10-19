package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://jatin.dev:3000/profile?skill=react-native&age=21"

func main() {
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	for _, val := range(qparams) {
		fmt.Println(val)
	}
}