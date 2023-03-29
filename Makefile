.PHONY: migrate migrate_down migrate_up migrate_version docker prod docker_delve local swaggo test
VERSION ?= $(shell git describe --tags --always)
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS ?= -X github.com/JamesHsu333/phantom_mask/pkg/version.Version=$(VERSION) -X github.com/JamesHsu333/phantom_mask/pkg/version.BuildDate=$(BUILD_DATE)

# Main
run:
	go run ./cmd/server/...

build:
	go build -ldflags="$(LDFLAGS)" ./cmd/server/...

# sqlc command
sql:
	sqlc generate -x

# Docker compose commands
up:
	echo "Starting production environment"
	docker-compose -f docker-compose.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build

# Go modules
tidy:
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

# Docker compose
FILES := $(shell docker ps -aq)

down:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

# Go migrate postgresql
force:
	migrate -database postgres://kdan:kdan@localhost:5432/kdan?sslmode=disable -source file://internal/data/migrations force 1

version:
	migrate -database postgres://kdan:kdan@localhost:5432/kdan?sslmode=disable -source file://internal/data/migrations version

migrate_up:
	migrate -database postgres://kdan:kdan@localhost:5432/kdan?sslmode=disable -source file://internal/data/migrations up

migrate_down:
	migrate -database postgres://kdan:kdan@localhost:5432/kdan?sslmode=disable -source file://internal/data/migrations down


