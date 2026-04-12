# Operations

## Docker Compose

From the repository root:

```bash
docker compose up --build
```

Services:

- `db`: PostgreSQL 16, credentials and database name match `docker-compose.yml`.
- `api`: built from the `Dockerfile`, listens on host port `8080` unless you change port mappings.

The API container waits for the database healthcheck before starting.

## Single binary image

The `Dockerfile` produces a static Linux `amd64` binary and runs it as a non-root user on distroless. Adjust `GOARCH` / `GOOS` in the build stage if you publish multi-arch images.

## Makefile

| Target | Action |
|--------|--------|
| `make tidy` | `go mod tidy` |
| `make test` | `go test -race ./...` |
| `make run` | `go run ./cmd/api` (requires `DATABASE_URL` in the environment) |

## Graceful shutdown

The process listens for `SIGINT` and `SIGTERM`. On shutdown it stops accepting new connections, waits for in-flight HTTP handlers up to `SHUTDOWN_TIMEOUT_MS`, then closes the database pool. Press **Ctrl+C** once in the terminal running the server and wait for shutdown to finish.

## Observability

Logs are structured JSON to stdout via the standard library `log/slog` package. Use your platform’s log aggregation (CloudWatch, Loki, Datadog agent, and so on) by shipping container stdout.

## Production checklist (beyond this repo)

This repository is a starter. Before exposing it on the public internet you would typically add authentication, authorization, rate limiting, TLS termination at a reverse proxy or load balancer, secret management, backups, and dependency and image scanning aligned with your organization’s policy.
