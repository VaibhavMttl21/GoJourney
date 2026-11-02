package main

import (
	"context"
	"log"
	"net/http"
	"uuid"

	"auth-jwt/internal/auth"
	"auth-jwt/internal/user"

	"github.com/gorilla/mux"
)

func main() {
	userRepo := user.NewInMemoryRepository()

	// Seed admin user
	adminHash, err := auth.HashPassword("admin123")
	if err != nil {
		log.Fatal(err)
	}

	_, err = userRepo.Create(context.Background(), user.User{
		ID:           uuid.New().String(),
		Email:        "admin@example.com",
		PasswordHash: adminHash,
		Role:         "admin",
	})
	if err != nil {
		log.Fatal(err)
	}

	refreshTokenStore := auth.NewRefreshTokenStore()

	authService := auth.NewService(
		userRepo,
		refreshTokenStore,
	)
	authHandler := auth.NewHandler(authService)

	r := mux.NewRouter()

	r.HandleFunc("/signup", authHandler.Signup).Methods("POST")
	r.HandleFunc("/signin", authHandler.Signin).Methods("POST")
	r.HandleFunc("/signout", authHandler.Signout).Methods("POST")
    r.HandleFunc("/refresh",authHandler.Refresh,).Methods("POST")

	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(auth.AuthMiddleWare)

	protected.HandleFunc("/profile", authHandler.Profile).Methods("GET")

	admin := protected.PathPrefix("/admin").Subrouter()
	admin.Use(auth.ReqRole("admin"))

	admin.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome Admin"))
	}).Methods("GET")

	log.Println("server running on :4000")

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}
