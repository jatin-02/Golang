package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"courseprice"`
	Author *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"website"`
}

// Middleware or helper
func (c *Course) isEmpty() bool {
	return c.CourseId=="" && c.CourseName=="" 
}

// Fake db
var courses []Course

func main() {
	fmt.Println("API")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId:"1", CourseName: "React", CoursePrice: 200, Author: &Author{FullName:"Jatin", Website:"jatin.dev"}})
	courses = append(courses, Course{CourseId:"2", CourseName: "React-Native", CoursePrice: 300, Author: &Author{FullName:"Jatin", Website:"udemy"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("Get")
	r.HandleFunc("/courses", getAllCourses).Methods("Get")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("Get")
	r.HandleFunc("/course/{id}", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen on port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controller
func serveHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to Home Page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	fmt.Println("Course not found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside Json")
		return
	}

	// rand.Seed(time.Now().UnixNano())
	// course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Course Updated")
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}