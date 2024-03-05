.PHONY: generateAndRun
generateAndRun:
	@./../../Go/bin/templ generate
	@cd frontend && go run cmd/*.go

.PHONE: migrateUp
migrateUp:
	@migrate -path=./migrations/ -database=${GOPOLL_DB_DSN}'?sslmode=disable' up

.PHONE: migrateDown
migrateDown:
	@migrate -path=./migrations/ -database=${GOPOLL_DB_DSN}'?sslmode=disable' down