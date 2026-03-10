.PHONY: up down build logs generate migrate-new lint test

up:
	docker compose up --build -d

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

migrate-new:
	atlas migrate diff $(name) --env local

migrate-apply:
	atlas migrate apply --env local

lint:
	golangci-lint run ./...

test:
	go test ./...
