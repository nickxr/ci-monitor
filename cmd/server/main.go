package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("CI Monitor is running!"))
	})

	log.Println("Server starting on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
