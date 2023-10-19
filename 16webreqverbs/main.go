package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello from Jatin")
	PerformGetRequest()
}

func PerformGetRequest()  {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	CheckNilError(err)
	defer response.Body.Close()

	fmt.Println("Status Code:", response.Status)
	fmt.Println("Content length:", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	CheckNilError(err)
	fmt.Println(string(content))
}

func PerformPostRequest()  {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"Name":"Jatin",
			"age":21,
			"status":true
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	CheckNilError(err)

	content, err := io.ReadAll(response.Body)
	CheckNilError(err)

	fmt.Println(string(content))
}

func PerformPostFormRequest()  {
	const myurl = "http://localhost:8000/post"

	data := url.Values{}
	data.Add("firstName", "Jatin")
	data.Add("lastName", "Soni")
	response, err := http.PostForm(myurl, data)
	CheckNilError(err)

	content, err := io.ReadAll(response.Body)
	CheckNilError(err)

	fmt.Println(string(content))
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}