// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/courses", getcourses).Methods("GET")

// 	r.Use(loggingmiddleware)

// 	log.Fatal(http.ListenAndServe(":4000", r))
// }

// func getcourses(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Inside getCourses")
// 	w.Write([]byte("Courses"))
// }

// func loggingmiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("in the middleware")
// 		fmt.Println("Method:", r.Method)
// 		fmt.Println("Path:", r.URL.Path)

// 		next.ServeHTTP(w, r)
// 	})
// }
// // auth middlweware with context
// func authMiddleware(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		authHeader := r.Header.Get("Authorization")

// 		if authHeader == "" {

// 			http.Error(
// 				w,
// 				"Authorization required",
// 				http.StatusUnauthorized,
// 			)

// 			return
// 		}

// 		// Later:
// 		// 1. Extract Bearer token
// 		// 2. Verify JWT signature
// 		// 3. Validate expiry
// 		// 4. Extract user ID
// 		// 5. Put user ID into Context

// 		userID := "123"

// 		ctx := context.WithValue(
// 			r.Context(),
// 			userIDKey,
// 			userID,
// 		)

// 		r = r.WithContext(ctx)

//			next.ServeHTTP(w, r)
//		})
//	}
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
	roleKey      contextKey = "role"
)

func main() {

	r := mux.NewRouter()

	// Global middleware
	r.Use(recoveryMiddleware)
	r.Use(requestIDMiddleware)
	r.Use(loggingMiddleware)
	r.Use(corsMiddleware)

	// Public routes
	r.HandleFunc("/courses", getCourses).Methods("GET")

	// Protected API
	api := r.PathPrefix("/api").Subrouter()

	api.Use(authMiddleware)
	// r.Handle("/profilee",authMiddleware(http.HandlerFunc(profileHandler))).Methods("GET")

	api.HandleFunc("/profile", profileHandler).Methods("GET")

	// Admin-only route
	admin := api.PathPrefix("/admin").Subrouter()

	admin.Use(adminMiddleware)

	admin.HandleFunc("/dashboard", adminDashboard).Methods("GET")

	fmt.Println("Server running on :4000")

	log.Fatal(http.ListenAndServe(":4000", r))
}
// ================ MIDDLEWARES ===================
func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		fmt.Println(
			"START",
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)

		fmt.Println(
			"END",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
func requestIDMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := uuid.New().String()

		ctx := context.WithValue(
			r.Context(),
			requestIDKey,
			requestID,
		)

		r = r.WithContext(ctx)

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r)
	})
}
func recoveryMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {

				fmt.Println("PANIC:", err)

				http.Error(
					w,
					"Internal Server Error",
					http.StatusInternalServerError,
				)
			}

		}()

		next.ServeHTTP(w, r)
	})
}
func corsMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set(
			"Access-Control-Allow-Origin",
			"*",
		)

		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS",
		)

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func authMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {

			http.Error(
				w,
				"Unauthorized",
				http.StatusUnauthorized,
			)

			return
		}

		// JWT verification goes here.

		userID := "123"
		role := "admin"

		ctx := context.WithValue( // creates a new child Context based on the request's existing Context.
			r.Context(),
			userIDKey,
			userID,
		)
		// You're creating a new child Context that 
		// derives from r.Context() and associates a key/value with it.

		ctx = context.WithValue(
			ctx,
			roleKey,
			role,
		)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
func adminMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		role, ok := r.Context().
			Value(roleKey).(string)

		if !ok || role != "admin" {

			http.Error(
				w,
				"Forbidden",
				http.StatusForbidden,
			)

			return
		}

		next.ServeHTTP(w, r)
	})
}
// ============ HANDLERS ========================
func getCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Public courses")
}
func profileHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(userIDKey)

	fmt.Fprintln(w, "Profile of user:", userID)
}
func adminDashboard(w http.ResponseWriter, r *http.Request){
	fmt.Println("This is admin dashboard ")
}
