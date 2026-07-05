package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nickxr/ci-monitor/internal/repository"
)

type GitHubWebhookPayload struct {
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Ref    string `json:"ref"`
	After  string `json:"after"`
	Status string `json:"status"`
}

func WebhookHandler(repo *repository.BuildRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p GitHubWebhookPayload
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if p.Repository.FullName == "" || p.Ref == "" || p.After == "" {
			http.Error(w, "missing fields", http.StatusBadRequest)
			return
		}

		branch := p.Ref
		if len(branch) > 11 && branch[:11] == "refs/heads/" {
			branch = branch[11:]
		}

		status := p.Status
		if status != "success" && status != "failure" {
			status = "pending"
		}

		id, err := repo.Create(r.Context(), p.Repository.FullName, branch, p.After, status)
		if err != nil {
			log.Printf("db error: %v\n", err)
			http.Error(w, "internal error", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":      id,
			"status":  status,
			"message": "build recorded",
		})
	}
}
