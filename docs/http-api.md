# HTTP API

Base URL in local examples: `http://localhost:8080`. All JSON responses use UTF-8 and `Content-Type: application/json` unless noted.

## Health

### `GET /healthz`

Liveness: process is up. No database check.

**Response:** `200` body `ok` (plain text).

### `GET /readyz`

Readiness: PostgreSQL reachable with `Ping` on the application pool.

| Status | Body |
|--------|------|
| `200` | `ready` (plain text) |
| `503` | `not_ready` (plain text) |

## Users (`/v1`)

### `POST /v1/users`

Creates a user. Request body is capped at 1 MiB.

**Request JSON**

| Field | Type | Notes |
|-------|------|--------|
| `email` | string | Trimmed and stored lowercased by the domain layer |
| `name` | string | Trimmed |

**Responses**

| Status | Meaning |
|--------|---------|
| `201` | Created; body is a user object |
| `400` | Invalid JSON or invalid input (`{"error":"invalid_json"}` or `invalid_input`) |
| `409` | Email already exists (`conflict`) |
| `500` | Unexpected error (`internal`) |

**Example**

```bash
curl -sS -X POST http://localhost:8080/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"email":"you@example.com","name":"You"}'
```

**Success body shape**

```json
{
  "id": "uuid",
  "email": "you@example.com",
  "name": "You",
  "created_at": "2026-04-12T12:00:00Z"
}
```

### `GET /v1/users/{id}`

`{id}` must be a UUID.

| Status | Meaning |
|--------|---------|
| `200` | User object as above |
| `400` | Malformed UUID (`invalid_id`) |
| `404` | Unknown id (`not_found`) |
| `500` | Unexpected error (`internal`) |

**Example**

```bash
curl -sS http://localhost:8080/v1/users/00000000-0000-0000-0000-000000000001
```
