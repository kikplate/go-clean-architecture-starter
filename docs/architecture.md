# Architecture

This starter is a single deployable HTTP service. Dependencies point inward: outer layers depend on inner abstractions, not the reverse.

## Layers

| Area | Path | Role |
|------|------|------|
| Entry | `cmd/api` | Process entry: logging, signal handling, config load, bootstrap |
| Composition | `internal/app` | Wires config, database, use cases, HTTP router, and server lifecycle |
| Delivery | `internal/delivery/http` | Chi routes, middleware, JSON request and response shapes |
| Use cases | `internal/usecase` | Application rules orchestrating domain and ports |
| Domain | `internal/domain` | Entities, validation, repository interfaces, domain errors |
| Persistence | `internal/repository/postgres` | PostgreSQL implementation of repository ports |
| Config | `internal/config` | Environment-backed settings |

## Request flow

A typical HTTP call is handled in this order:

1. Chi middleware (request id, real client IP, panic recovery, per-request timeout).
2. HTTP handler decodes input, maps errors to status codes.
3. Use case executes rules using domain types and calls a repository port.
4. Repository runs SQL via `pgxpool`.
5. Handler encodes JSON or an error envelope.

## Migrations

Schema is applied at startup in `internal/repository/postgres/migrate.go` using idempotent DDL (`create table if not exists`). For larger products you would replace this with a dedicated migration tool and versioned SQL files.

## Extending the template

Add new domain types and ports in `internal/domain`, new use case packages under `internal/usecase`, new repository methods or tables under `internal/repository/postgres`, and register routes in `internal/delivery/http`. Keep `cmd/api` thin so composition stays in `internal/app`.
