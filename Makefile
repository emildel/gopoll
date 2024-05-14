.PHONY: generateAndRun
generateAndRun:
	@./../../Go/bin/templ generate
	@go run cmd/*.go

## migrateNew name=$1: create a new database migration
.PHONY: migrateNew
migrateNew:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: migrateUp
migrateUp:
	@migrate -path=./migrations/ -database=${GOPOLL_DB_DSN}'?sslmode=disable' up

.PHONY: migrateDown
migrateDown:
	@migrate -path=./migrations/ -database=${GOPOLL_DB_DSN}'?sslmode=disable' down

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/gopoll: build the cmd/api application
.PHONY: build/gopoll
build/gopoll:
	@echo 'Building gopoll...'
	@go build -ldflags='-s -w' -o=./bin/gopoll ./cmd