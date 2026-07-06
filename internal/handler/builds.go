package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nickxr/ci-monitor/internal/repository"
)

func BuildHandler(repo *repository.BuildRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		builds, err := repo.GetAll(r.Context())
		if err != nil {
			http.Error(w, "failed to fetch builds", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(builds)
	}
}

func BuildByIDHandler(repo *repository.BuildRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		_, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "not implemented yet",
			"id":      idStr,
		})
	}
}

func StatsHandler(repo *repository.BuildRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		builds, err := repo.GetAll(r.Context())
		if err != nil {
			http.Error(w, "failed to fetch builds", http.StatusInternalServerError)
			return
		}

		var pending, success, failure int
		for _, b := range builds {
			switch b.Status {
			case "pending":
				pending++
			case "success":
				success++
			case "failure":
				failure++
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"total":   len(builds),
			"pending": pending,
			"success": success,
			"failure": failure,
		})
	}
}
