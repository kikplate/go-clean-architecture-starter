# go-clean-architecture-starter

Go HTTP API starter using Chi, PostgreSQL, and layered boundaries between delivery, use cases, domain, and persistence.

## Layout

- `cmd/api` service entrypoint
- `internal/delivery/http` routing and HTTP adapters
- `internal/usecase` application workflows
- `internal/domain` entities and repository ports
- `internal/repository/postgres` PostgreSQL adapters
- `internal/config` environment-backed configuration
- `internal/app` composition and server lifecycle

## Run locally

```bash
docker compose up --build
```

## Configure

Copy `.env.example` to `.env` and adjust values. Required environment variable:

- `DATABASE_URL`

Optional:

- `HTTP_ADDR` (default `:8080`)
- `LOG_LEVEL` (`info` or `debug`)
- `REQUEST_TIMEOUT_MS`
- `SHUTDOWN_TIMEOUT_MS`

## HTTP surface

- `GET /healthz`
- `GET /readyz`
- `POST /v1/users` JSON body `{"email":"...","name":"..."}`
- `GET /v1/users/{id}`

## Tests

```bash
make test
```

Fast tests cover configuration parsing, domain normalization, use cases, HTTP handlers, and routing behavior. PostgreSQL repository tests run when `DATABASE_URL` is set (for example in CI or with `docker compose up -d db` and a local DSN). Without `DATABASE_URL`, those tests are skipped so `go test ./...` stays usable on laptops without a database.

## Kikplate manifest

This repository includes `kikplate.yaml` for registry metadata. Set `owner` to your GitHub username before submitting the plate. After submission, add `verification_token` from Kikplate and run verification.
