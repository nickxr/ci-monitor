# CI Monitor Service

A lightweight REST API service that receives GitHub webhook events and stores build history in PostgreSQL.

## Features
- POST `/webhook` – receive GitHub webhook events.
- GET `/builds` – list all builds.
- GET `/builds/{id}` – get a single build.
- GET `/stats` – build statistics (pending/success/failure).

## Tech Stack
- Go 1.25 (chi router, pgx)
- PostgreSQL 15 (via Docker)
- Docker Compose
- GitHub Actions (for demo webhook delivery)

## Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/nickxr/ci-monitor.git
   cd ci-monitor

2. Start PostgreSQL:
   ```bash
   docker compose up -d db

3. Run migrations:
   ```bash
   make migrate

5. Start the server:
   ```bash
   make run
6. Send a test webhook:
   ```bash
   curl -X POST http://localhost:8080/webhook \
   -H "Content-Type: application/json" \
   -d '{
    "repository": {"full_name": "nickxr/ci-monitor"},
    "ref": "refs/heads/main",
    "after": "abc123def456",
    "status": "success"
   }'
  
## API Endpoints

| Method |  Endpoint   | Description |
|:-------|:-----------:| ----: |
| POST   |  /webhook   | Receive a GitHub webhook event |
| GET    |   /builds   | List all builds |
| GET    | /builds/{id} | Get a specific build by ID |
| GET    |    /stats    | Get build statistics |

## Testing
   
```bash
   go test -v ./...
