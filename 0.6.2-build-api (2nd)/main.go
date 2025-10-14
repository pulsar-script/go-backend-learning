package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// custome data types

// course data type
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"Author"`
}

// author data type
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var Courses []Course

//* middleware, helper - file

// IsEmpty
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

	fmt.Println("API - asimmomin")
	r := mux.NewRouter()

	//seeding
	Courses = append(Courses,
		Course{
			CourseId:    "123",
			CourseName:  "React JS tutorial",
			CoursePrice: 599,
			Author: &Author{
				Fullname: "Josef Maria",
				Website:  "www.JM.com",
			},
		},
		Course{
			CourseId:    "456",
			CourseName:  "Java Tutorial",
			CoursePrice: 799,
			Author: &Author{
				Fullname: "San Karla",
				Website:  "www.SK.uk",
			},
		},
	)

	// routing - file
	r.HandleFunc("/", sreverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//* controllers - file

// server home route

func sreverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API by Hero Corporation </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab if from request
	params := mux.Vars(r)

	// loop  through Courses, find nothing id and return the response
	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with give id") //TODO add id number here
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is nil, means not send
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if: body is empty,{}

	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)

	if newCourse.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO : If course is duplicate

	// generate unique id, string
	randomeGenerator := rand.New(rand.NewSource(time.Now().UnixMilli()))
	newCourse.CourseId = strconv.Itoa(randomeGenerator.Intn(100))

	// append course into Courses
	Courses = append(Courses, newCourse)
	json.NewEncoder(w).Encode(newCourse)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// grab all new coming data & id also
	params := mux.Vars(r)

	// find that course index from Courses []Course

	for index, course := range Courses {

		if course.CourseId == params["id"] {

			// remove old course which matches the id
			Courses = append(Courses[:index], Courses[index+1:]...)

			// create new course with new values (here we are not changing specific fields of values )
			var updateCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updateCourse)

			// insert id again
			updateCourse.CourseId = params["id"]

			// add updateCourse into []Course
			Courses = append(Courses, updateCourse)

			// send with response (w)
			json.NewEncoder(w).Encode(updateCourse)

			return

		}
	}

	// and case when id not found 404
	json.NewEncoder(w).Encode("No id found 404")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// grab id of course
	params := mux.Vars(r)

	// loop over
	for index, course := range Courses {

		// if : id exits
		if course.CourseId == params["id"] {

			// find index and remove from []Courses
			Courses = append(Courses[:index], Courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is successfully deleted")
			return
		}
	}

	// else : id no exists
	json.NewEncoder(w).Encode("Course with give id is not found") //TODO: fix this using label
	return
}
