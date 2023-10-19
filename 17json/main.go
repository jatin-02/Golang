package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Nmae     string `json:"CourseName"`
	Price    int
	Platform string
	Password string `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Slice to Json & Vice-versa")
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React", 200, "udemy", "abc123", []string{"frontend", "react"}},
		{"React Native", 300, "coursera", "ab3123", []string{"app", "react-native"}},
		{"Angular", 100, "udemy", "abc223", nil},
	}

	fileJson, err := json.MarshalIndent(courses, "", "\t")
	checkNilError(err)

	fmt.Printf("%s\n", fileJson)
}

func DecodeJson() {
	courses := []byte(`
		{
			"CourseName": "React",
			"Price": 200,
			"Platform": "udemy",
			"tags": ["frontend","react"]
        }
	`)

	var storeCourses course
	checkValid := json.Valid(courses)

	if checkValid {
		json.Unmarshal(courses, &storeCourses)
		fmt.Printf("%#v\n", storeCourses)
	} else {
		fmt.Println("JSON DATA IS INVALID")
	}
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}