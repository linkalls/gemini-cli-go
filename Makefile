# Makefile for gemini-cli (Go version)

.PHONY: help install build test lint format clean run

help:
	@echo "Makefile for gemini-cli (Go version)"
	@echo ""
	@echo "Usage:"
	@echo "  make install          - Install Go dependencies"
	@echo "  make build            - Build the Go binary"
	@echo "  make test             - Run the test suite"
	@echo "  make lint             - Lint the code (requires golangci-lint)"
	@echo "  make format           - Format the code"
	@echo "  make clean            - Remove generated files"
	@echo "  make run              - Run the Gemini CLI"

install:
	go mod download
	go mod tidy

build:
	go build -o gemini .

test:
	go test -v ./...

lint:
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	golangci-lint run

format:
	go fmt ./...
	gofmt -s -w .

clean:
	rm -f gemini
	go clean

run: build
	./gemini
