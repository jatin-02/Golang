package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["RJ"]= "ReactJS"
	languages["JS"]= "JavaScript"
	languages["RN"]= "React Native"
	fmt.Println(languages)
	fmt.Println(languages["JS"])
	
	delete(languages, "JS")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For the key %v, the value is %v\n", key, value)
	}
}