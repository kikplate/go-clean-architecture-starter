.PHONY: tidy test run docker-up docker-down

tidy:
	go mod tidy

test:
	go test -race ./...

run:
	go run ./cmd/api

docker-up:
	docker compose up --build

docker-down:
	docker compose down -v
