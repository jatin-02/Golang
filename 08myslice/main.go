package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"Apple", "Orange", "Tomato"}
	fmt.Println(fruits)
	
	fruits = append(fruits, "Banana", "Grapes")
	fmt.Println(fruits)
	
	fruits = append(fruits[1:3])
	fmt.Println(fruits)

	vege := make([]int, 3)
	vege[0] = 5
	vege[1] = 10
	vege[2] = 15
	fmt.Println(vege)

	vege = append(vege, 20, 25)
	fmt.Println(vege)

	sort.Ints(vege)
	fmt.Println(sort.IntsAreSorted(vege))

	var courses = []string{"React", "Next", "Swift", "Kotlin", "Golang"}
	fmt.Println(courses)
	
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}