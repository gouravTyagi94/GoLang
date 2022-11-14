// Medium blog to completely understand gorilla mux vs built-in http module
// https://levelup.gitconnected.com/experiment-golang-http-builtin-and-related-popular-packages-1d9a6dcb80d

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

//Model for course - file
type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"-"`

	//Actually, we can not refer to Author struct as it is not created yet
	//But this is a symbolic representation that this field is of type Author
	//Hence, it's written as *Author
	Author *Author `json:""author`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db
var courses []Course

//Middleware or Helper func-- file
func (c *Course) IsEmpty() bool {
	// return c.CourseId=="" && c.CourseName==""
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	//seed the data
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299,
		Author: &Author{Fullname: "Gourav", Website: "gourav.dev"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199,
		Author: &Author{Fullname: "Gourav", Website: "gourav.dev"}})

	
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")

	//get all courses
	r.HandleFunc("/courses", getAllCourses).Methods("GET")

	// This /{id} which we are giving along with URL will be captured in params
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")

	//create a course
	r.HandleFunc("/course", createOneCourse).Methods("POST")

	//Update one course
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")

	//delete a course
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	//listening on a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers-- file

//serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to our Golang Course</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-type", "application/json")

	//What if -- body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about  - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	//generate unique id and convert it into string
	//add courses into courses slice

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-type", "application/json")

	// mux.Vars(r) returns all values captured in the request URL.
	// vars is a dictionary whose key-value pairs are variables' name-value pairs

	//first grab id from req
	params := mux.Vars(r)

	//loop, get id, remove, add with my id

	for index, course := range courses {
		//get id
		if course.CourseId == params["id"] {
			//removing
			courses = append(courses[:index], courses[index+1:]...)

			//updating with new data
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			//writing the response in the writer w
			json.NewEncoder(w).Encode(courses)
		}
	}
	//Todo
	//Send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Successfully removed the course...")
			break
		}
	}
}
