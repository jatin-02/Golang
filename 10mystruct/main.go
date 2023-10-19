package main

import "fmt"

func main() {
	Jatin := User{"Jatin Soni", "jatin1202@gmail.com", true, 21}
	fmt.Println(Jatin)
	fmt.Printf("User details are %+v\n", Jatin)
	fmt.Printf("User name is %v\n", Jatin.Name)

	Jatin.GetStatus()
	Jatin.UpdateEmail()
	fmt.Println("Does Email update:", Jatin.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) UpdateEmail() {
	u.Email = "Jatin1202new@gmail.com"
	fmt.Println("Method call for Updated Email:", u.Email)
}