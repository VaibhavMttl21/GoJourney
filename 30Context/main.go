// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func Worker(ctx context.Context) {
// 	for i := 0; i < 10; i++ {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("worker finished")
// 			return
// 		default:
// 			fmt.Println("worker working with id", i)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go Worker(ctx)

// 	time.Sleep(3* time.Second)

// 	cancel()

// 	time.Sleep(1* time.Second)

// }
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/slow", slowHandler).Methods("GET")

	fmt.Println("Server running on :4000")

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func slowHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")

	// Get the Context associated with this HTTP request
	ctx := r.Context()

	// Create a child Context with a 3-second timeout
	ctx, cancel := context.WithTimeout(
		ctx,
		6*time.Second,
	)

	defer cancel()

	fmt.Println("Request started")

	// Call fake database
	course, err := getCourseFromDB(ctx)

	if err != nil {
		fmt.Println("There is some error")
		return
	}

	fmt.Fprintln(w, course)
}

func getCourseFromDB(ctx context.Context) (string, error) {

	fmt.Println("Database operation started")

	select {

	case <-time.After(5 * time.Second):

		fmt.Println("Database operation completed")

		return "Golang Backend Development", nil

	case <-ctx.Done():

		fmt.Println("Database operation cancelled")

		return "", ctx.Err()
	}
}