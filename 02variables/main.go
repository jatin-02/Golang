package main

import "fmt"

const LoginToken = 120324 //Public

func main() {
	var username string = "Jatin"
	fmt.Println(username)
	fmt.Printf("username is of type: %T \n", username)
	
	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("username is of type: %T \n", isLoggedin)

	simple := "But not outside"
	fmt.Println(simple)

	fmt.Println(LoginToken)
}