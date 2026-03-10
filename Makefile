.PHONY: up down build logs generate migrate migrate-apply migrate-check lint test

up:
	docker compose up --build -d

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

# Generate Ent code + auto-create migration from schema diff
generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

migrate:
	@if [ -z "$(name)" ]; then echo "Usage: make migrate name=<migration_name>"; exit 1; fi
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
	atlas migrate diff $(name) --env local
	atlas migrate lint --env local --latest 1

# Check that migrations are up to date with Ent schema (used in CI)
migrate-check:
	atlas migrate diff --env local check-sync

migrate-apply:
	atlas migrate apply --env local

lint:
	golangci-lint run ./...

test:
	go test ./...
