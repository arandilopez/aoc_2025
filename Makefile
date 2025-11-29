.PHONY: run build test

# Makefile for building the AoC application
# Usage:
#   make build  - to build the command-line application
#   make clean  - to clean up the build artifacts
#   make test   - to run the tests

build:
	@go build -o bin/aoc cmd/aoc/main.go

clean:
	@rm -rf bin/aoc

test:
	@go test ./test/...
