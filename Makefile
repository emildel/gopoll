.PHONY: generateAndRun
generateAndRun:
	@./../../Go/bin/templ generate
	@cd frontend && go run cmd/*.go

## migrateNew name=$1: create a new database migration
.PHONY: migrateNew
migrateNew:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONE: migrateUp
migrateUp:
	@migrate -path=./migrations/ -database=${GOPOLL_DB_DSN}'?sslmode=disable' up

.PHONE: migrateDown
migrateDown:
	@migrate -path=./migrations/ -database=${GOPOLL_DB_DSN}'?sslmode=disable' down