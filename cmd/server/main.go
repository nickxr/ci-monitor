package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nickxr/ci-monitor/internal/config"
	"github.com/nickxr/ci-monitor/internal/db"
	"github.com/nickxr/ci-monitor/internal/handler"
	"github.com/nickxr/ci-monitor/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("config load:", err)
	}

	dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"
	dbPool, err := db.NewDB(dsn)
	if err != nil {
		log.Fatal("db connect:", err)
	}
	defer dbPool.Close()

	repo := repository.NewBuildRepository(dbPool)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("CI Monitor is running!"))
	})

	r.Post("/webhook", handler.WebhookHandler(repo))

	log.Println("starting server on :" + cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
