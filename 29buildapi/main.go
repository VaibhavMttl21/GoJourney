package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"uuid"

	"github.com/gorilla/mux"
)

// model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// helper -file

func (c *Course) IsEmpty() bool {
	// return c.CourseName == "" && c.CourseId == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Api in go lang")
	r := mux.NewRouter()
	courses = append(courses,
		Course{
			CourseId:    "3",
			CourseName:  "Golang Backend Development",
			CoursePrice: 499,
			Author: &Author{
				Fullname: "Rahul Sharma",
				Website:  "rahul.dev",
			},
		},
		Course{
			CourseId:    "4",
			CourseName:  "Node.js and Express",
			CoursePrice: 399,
			Author: &Author{
				Fullname: "Aman Verma",
				Website:  "aman.dev",
			},
		},
		Course{
			CourseId:    "5",
			CourseName:  "React and TypeScript",
			CoursePrice: 599,
			Author: &Author{
				Fullname: "Priya Singh",
				Website:  "priya.dev",
			},
		},
		Course{
			CourseId:    "6",
			CourseName:  "PostgreSQL for Developers",
			CoursePrice: 299,
			Author: &Author{
				Fullname: "Arjun Mehta",
				Website:  "arjun.dev",
			},
		},
		Course{
			CourseId:    "7",
			CourseName:  "Docker and Kubernetes",
			CoursePrice: 699,
			Author: &Author{
				Fullname: "Rohit Kapoor",
				Website:  "rohit.dev",
			},
		},
		Course{
			CourseId:    "8",
			CourseName:  "System Design Fundamentals",
			CoursePrice: 799,
			Author: &Author{
				Fullname: "Neha Gupta",
				Website:  "neha.dev",
			},
		},
	)
	// routing
	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/courses", getallcourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOnecourse).Methods("GET")
	r.HandleFunc("/course", createOnecourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOnecourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOnecourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers -file
//servehome route

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome to serve home route </h1>"))
}

func getallcourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println(courses)
	w.Header().Set("Content-Type", "application/json")
	// we can do it through marshal also but this encoder is cleaner
	//   data, err := json.Marshal(courses)
	// if err != nil {
	//     http.Error(w, "Unable to encode courses", http.StatusInternalServerError)
	//     return
	// }

	// w.Write(data)
	json.NewEncoder(w).Encode(courses)
}

func getOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	fmt.Printf("the value of params is %v and the type of params is %T\n", params, params)

	//loop through courses and find id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given id")
	// return

}

func createOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data")
		return
	}

	// what abt data like this {}
	// if r.Body == "{}" {
	// 	json.NewEncoder(w).Encode("No data 2")
	// 	return
	// }
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	// fmt.Println("value is: ",value)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No course id and course name")
		return
	}
	// TODO check only title name matches
	for _, v := range courses {
		if v.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("course name is same")
			return
		}
	}
	// generate unique id , string
	// append course into courses

	// rand.Seed(time.Now().UnixNano())
	id := uuid.New().String()
	course.CourseId = id
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	// return
}

func updateOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from req
	params := mux.Vars(r)

	// loop,id,remove,add with my id

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("New Entry is added")
			return
		}
	}
	json.NewEncoder(w).Encode("Id is not found or data is empty")
	// if id is not found

}

func deleteOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode(
				fmt.Sprintf("that this id %s is removed", params["id"]),
			)
			break
		}
	}
}
