# Run the application locally.
run:
	go run cmd/server/main.go

# Apply all pending database migrations.
migrate:
	@export $$(cat .env | xargs) && migrate -path migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up

# Rollback the last migration.
migrate-down:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

# Run all unit tests.
test:
	go test -v ./...

# Build a binary for the current OS.
build:
	go build -o bin/ci-monitor cmd/server/main.go

# Build the Docker image.
docker-build:
	docker build -t ci-monitor .

# Start all Docker containers in detached mode.
docker-up:
	docker-compose up -d