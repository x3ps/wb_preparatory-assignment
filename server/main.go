package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	InitDB()
	defer DB.Close()

	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// endpoints
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	r.Get("/search", SearchCarsHandler)

	log.Println("Server running at http://0.0.0.0:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
