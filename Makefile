.PHONY: generateAndRun
generateAndRun:
	./../../Go/bin/templ generate
	cd frontend && go run cmd/*.go
