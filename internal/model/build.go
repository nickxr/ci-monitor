package model

import "time"

type BuildStatus string

const (
	StatusPending BuildStatus = "pending"
	StatusSuccess BuildStatus = "success"
	StatusFailure BuildStatus = "failure"
)

// CreatedAt and UpdatedAt are updated automatically in the database
type Build struct {
	ID         int64       `json:"id"`
	Repo       string      `json:"repo"`
	Branch     string      `json:"branch"`
	CommitHash string      `json:"commit_hash"`
	Status     BuildStatus `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
