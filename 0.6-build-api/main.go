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

// Model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // it is like 	REFERENCE KEY , we neet to give Author is pointer which point to Author struct (which ever author we create using Author struct)
	// &Author means we are giving memory address of already created author using Author struct, that dont work here , that is we giving value , here we want to give type
}

// Model for author - file
type Author struct {
	FullName string `json:"authorname"`
	Website  string `json:"website"`
}

// fake DB - slice of Course struct
var Courses []Course

// middleware, helper, method - file

func (c *Course) InEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

	fmt.Println("welcome to Udemy Clone")

	//* initializing NewRouter
	r := mux.NewRouter()

	// seeding - activity of inserting some dummy data into DB
	Courses = append(Courses,
		Course{
			CourseId:    "2",
			CourseName:  "NodeJS Web Development Course",
			CoursePrice: 344,
			Author: &Author{
				FullName: "Hitesh Choudary",
				Website:  "hcdev.com",
			},
		},
		Course{
			CourseId:    "4",
			CourseName:  "React Frontend Course",
			CoursePrice: 544,
			Author: &Author{
				FullName: " Amit Padman",
				Website:  "apdev.com",
			},
		},
	)

	//* routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	// /course/{id}  <= the name you give in { ... } , you have to use same while accessing it
	//  params["id"]
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//* listing to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// Controllers - files

// w => it use to write in response we will send
// r => it use to read request which come from frontend

// serve home route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API by asim | home route </h1>"))
}

// get all Courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

// get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	// loop through Courses, find matching id and return the response
	for _, course := range Courses {
		if course.CourseId == params["id"] { // we will learn about "id" parameter
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return
}

// create one course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// case handling

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if: body is send but json is empty
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course) // decode and where to store json data , here we are using METHOD 1 storing into data-structure form

	if course.InEmpty() {
		json.NewEncoder(w).Encode("No data in JSON")
	}

	// have data

	//TODO generate unique id & convert into string
	//TODO append course into Courses

	// 1. Create a new random source based on the current time in nanoseconds.
	//    This provides the unique seed.

	// 2. Create a new random generator from that source.
	randNumbGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 3. Use the new generator to get a random number.
	//    For example, a number between 0 and 99.
	course.CourseId = strconv.Itoa(randNumbGen.Intn(100))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// update one course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//TODO get and store id from request (params)
	params := mux.Vars(r)

	//TODO loop through Courses (kind a find course which matches id)
	for index, course := range Courses {
		if course.CourseId == params["id"] { // we will discuss this "id" while learning routing

			//TODO when find course remove it from DB ( Not remove in actule DB )
			Courses = append(Courses[:index], Courses[index+1:]...)

			//TODO take new updated data for that course from request body and store into local var
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse) // we do decoding when we get json data from request

			//TODO Add new updated course into Courses with id get from params
			course.CourseId = params["id"]           // we are just overwritting same value, for assurance
			Courses = append(Courses, updatedCourse) // adding updated course into Courses
			// our main work is to just update course into DB

			//TODO send some response to frontend , like "Successfully update" or we can send updated course
			json.NewEncoder(w).Encode(updatedCourse)

			return

		}
	}

}

// delete one course

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	//TODO get and store id of course user want to delete
	params := mux.Vars(r)

	//TODO find course from DB which match id
	for index, course := range Courses {
		if course.CourseId == params["id"] {

			//TODO remove course from Courses
			Courses = append(Courses[:index], Courses[index+1:]...)

			//TODO send some response to frontend , like "Successfully done"
			json.NewEncoder(w).Encode("Succssfully delete")
			break
		}
	}

}
