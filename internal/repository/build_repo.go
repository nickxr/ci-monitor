package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickxr/ci-monitor/internal/model"
)

type BuildRepository struct {
	db *pgxpool.Pool
}

func NewBuildRepository(db *pgxpool.Pool) *BuildRepository {
	return &BuildRepository{db: db}
}

func (r *BuildRepository) Create(ctx context.Context, repo, branch, commitHash, status string) (int64, error) {
	var id int64
	query := `
		INSERT INTO builds(repo, branch, commit_hash, status, created_at, updated_at)
		VALUES ($1,$2,$3,$4, NOW(), NOW())
		RETURNING id
		`
	err := r.db.QueryRow(ctx, query, repo, branch, commitHash, status).Scan(&id)
	return id, err
}

func (r *BuildRepository) GetAll(ctx context.Context) ([]model.Build, error) {
	query := `SELECT id, repo, branch, commit_hash, status, created_at, updated_at FROM builds ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var builds []model.Build
	for rows.Next() {
		var b model.Build
		if err := rows.Scan(&b.ID, &b.Repo, &b.Branch, &b.CommitHash, &b.Status, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		builds = append(builds, b)
	}
	return builds, nil
}
