DB_URL=postgres://postgres:postgres@localhost:5433/mining_asset?sslmode=disable
APP_NAME=mining-asset-management
DOCKER_IMAGE=afiffaizun/$(APP_NAME)

.PHONY: run lint test migrate-up migrate-down migrate-create docker-build docker-run

run:
	go run cmd/server/main.go

lint:
	golangci-lint run ./...

test:
	go test ./... -v -race

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir migrations -seq -digits 6 $$name

docker-build:
	docker build -t $(DOCKER_IMAGE):latest .

docker-run:
	docker compose up -d
