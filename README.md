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
3. Run migrations:
4. Start the server:
5. Send a test webhook:

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
