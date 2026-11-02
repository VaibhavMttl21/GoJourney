package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	database "course-api/internal/databse"
	"course-api/internal/handler"
	// "course-api/internal/repository/memory"
	"course-api/internal/repository/postgres"
	"course-api/internal/service"
)

func main() {
	db ,err := database.Connect()

	if err != nil{
		log.Fatal(err)
	}


	repo := postgres.NewPostgresCourseRepository(db)
	// Create repository implementation
	// repo := memory.NewInMemoryCourseRepository()

	// Inject repository into service
	courseService := service.NewCourseService(repo)

	// Inject service into handler
	courseHandler := handler.NewCourseHandler(courseService)

	fmt.Printf("%T\n",repo)
	fmt.Printf("%T\n",courseService)
	fmt.Printf("%T\n",courseHandler)
	// Create router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/course",courseHandler.CreateCourse).Methods("POST")
	r.HandleFunc("/courses",courseHandler.GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",courseHandler.GetByID).Methods("GET")
	r.HandleFunc("/course/{id}",courseHandler.ModifyCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",courseHandler.DelCourse).Methods("DELETE")
	// Start server
	log.Println("Server running on :4000")

	log.Fatal(
		http.ListenAndServe(":4000", r),
	)
}

// // when program starts
// // 					 main()
//     //                   │
//     //                   ↓
//     //     NewInMemoryCourseRepository()
//     //                   │
//     //                   ↓
//     //       InMemoryCourseRepository
//     //                   │
//     //                   ↓
//     //       NewCourseService(repo)
//     //                   │
//     //                   ↓
//     //            CourseService
//     //                   │
//     //                   ↓
//     //    NewCourseHandler(courseService)
//     //                   │
//     //                   ↓
//     //           CourseHandler
//     //                   │
//     //                   ↓
//     //           Register route
//     //                   │
//     //                   ↓
//     //           Start HTTP server							when request arrive

// 	                                                      // POST /course
// 														//       │
// 														//       ↓
// 														//    Router
// 														//       │
// 														//       ↓
// 														// CourseHandler.CreateCourse()
// 														//       │
// 														//       ↓
// 														// CourseService.Create()
// 														//       │
// 														//       ↓
// 														// repository.CourseRepository
// 														//       │
// 														//       ↓
// 														// InMemoryCourseRepository.Create()
// 														//       │
// 														//       ↓
// 														//    []Course

// package main

// import (
// 	database "course-api/internal/databse"
// 	"log"
// )

// func main() {

//     db, err := database.Connect()

//     if err != nil {
//         log.Fatal(err)
//     }

//     log.Println("Connected to PostgreSQL:", db != nil)
// }