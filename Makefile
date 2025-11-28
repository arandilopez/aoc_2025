.PHONY: run build

run:
	@go run cmd/aoc/main.go $(ARGS)

build:
	@go build -o bin/aoc cmd/aoc/main.go
