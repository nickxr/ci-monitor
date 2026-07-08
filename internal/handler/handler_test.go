package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickxr/ci-monitor/internal/model"
	"github.com/nickxr/ci-monitor/internal/repository"
	"github.com/stretchr/testify/assert"
)

func setupTestRepo(t *testing.T) *repository.BuildRepository {
	dsn := "postgres://postgres:postgres@localhost:5432/ci-monitor?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatalf("failed to connect to DB: %v", err)
	}
	return repository.NewBuildRepository(pool)
}

func TestWebhookHandler_Success(t *testing.T) {
	repo := setupTestRepo(t)
	r := chi.NewRouter()
	r.Post("/webhook", WebhookHandler(repo))

	payload := map[string]interface{}{
		"repository": map[string]string{"full_name": "test/repo"},
		"ref":        "refs/heads/main",
		"after":      "abc123",
		"status":     "success",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "ok", resp["status"])
	assert.NotNil(t, resp["id"])
}

func TestBuildsHandler_ReturnsList(t *testing.T) {
	repo := setupTestRepo(t)
	r := chi.NewRouter()
	r.Get("/builds", BuildHandler(repo))

	req := httptest.NewRequest(http.MethodGet, "/builds", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var builds []model.Build
	err := json.Unmarshal(w.Body.Bytes(), &builds)
	assert.NoError(t, err)
}
