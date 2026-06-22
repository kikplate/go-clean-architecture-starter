# Configuration

All settings are read from the process environment. There is no built-in `.env` loader: if you keep secrets or local overrides in a `.env` file, export variables into the shell before `go run` or `make run`, or rely on `docker compose` `environment` / `env_file` as in `docker-compose.yml`.

## Variables

| Name | Required | Default | Purpose |
|------|----------|---------|---------|
| `DATABASE_URL` | Yes | none | PostgreSQL DSN for `pgxpool` (include `sslmode` as required by your host) |
| `HTTP_ADDR` | No | `:8080` | Listen address passed to `http.Server` |
| `LOG_LEVEL` | No | `info` | `info` or `debug` for JSON log verbosity |
| `REQUEST_TIMEOUT_MS` | No | `15000` | Chi request timeout middleware duration |
| `SHUTDOWN_TIMEOUT_MS` | No | `20000` | Maximum time for graceful `http.Server.Shutdown` |

Invalid or non-positive millisecond values for the timeout variables fall back to the defaults above.

## Local examples

Docker Compose sets `DATABASE_URL` for the `api` service. For a shell on the host talking to Compose Postgres on port `5432`, a typical value matches `.env.example`.

Do not commit real credentials. `.env` is listed in `.gitignore`; use `.env.example` as a template only.
