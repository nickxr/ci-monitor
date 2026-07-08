package repository

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {
	dsn := "postgres://postgres:postgres@localhost:5432/ci_monitor?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatalf("failed to connect to test DB: %v", err)
	}
	return pool
}

func TestCreateAndGetAll(t *testing.T) {
	pool := setupTestDB(t)
	defer pool.Close()

	repo := NewBuildRepository(pool)

	id, err := repo.Create(context.Background(), "test/repo", "main", "abc123", "success")
	assert.NoError(t, err)
	assert.Greater(t, id, int64(0))

	builds, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, builds)

	var found bool
	for _, b := range builds {
		if b.ID == id {
			found = true
			assert.Equal(t, b.Repo, "test/repo")
			assert.Equal(t, b.Branch, "main")
			assert.Equal(t, b.CommitHash, "abc123")
			assert.Equal(t, b.Status, "success")
			break
		}
	}
	assert.True(t, found, "created build should be found in GetAll")
}

func TestCreateAndGetByID(t *testing.T) {
	pool := setupTestDB(t)
	defer pool.Close()

	repo := NewBuildRepository(pool)

	id, err := repo.Create(context.Background(), "test/repo", "main", "def456", "failure")
	assert.NoError(t, err)

	build, err := repo.GetByID(context.Background(), id)
	assert.NoError(t, err)
	assert.NotNil(t, build)
	assert.Equal(t, id, build.ID)
	assert.Equal(t, build.Repo, "test/repo")
	assert.Equal(t, string(build.Status), "failure")
}
