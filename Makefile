.PHONY: run build

# Makefile for building the AoC application
# Usage:
#   make build  - to build the command-line application

build:
	@go build -o bin/aoc cmd/aoc/main.go
