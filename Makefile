.PHONY: lint lint-fix lint-fast build run clean test test-verbose test-cover test-cover-html check help

# Lint
lint:
	golangci-lint run --config ./golangci.yml

lint-fix:
	golangci-lint run --fix --config ./golangci.yml

lint-fast:
	golangci-lint run --fast --config ./golangci.yml

# Build
build:
	@echo "Building application..."
	@mkdir -p bin
	@go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
	@echo "Build complete: bin/hexlet-path-size"

# Run
run:
	@go run ./cmd/hexlet-path-size

# Clean
clean:
	@rm -rf bin
	@echo "Cleaned up bin directory"

# Tests
test:
	@echo "Running tests..."
	go test ./...

test-verbose:
	@echo "Running tests with verbose output..."
	go test -v ./...

test-cover:
	@echo "Running tests with coverage..."
	go test -cover ./...

test-cover-html:
	@echo "Running tests with coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Check
check: test build
	@echo "All checks passed!"

# Help
help:
	@echo "Available commands:"
	@echo "  make lint           - run linter"
	@echo "  make lint-fix       - run linter with auto-fix"
	@echo "  make lint-fast      - run linter fast mode"
	@echo "  make build          - build the binary"
	@echo "  make run            - run the application"
	@echo "  make clean          - clean build artifacts"
	@echo "  make test           - run tests"
	@echo "  make test-verbose   - run tests with verbose output"
	@echo "  make test-cover     - run tests with coverage"
	@echo "  make test-cover-html - generate coverage report"
	@echo "  make check          - run tests and build"
	@echo "  make help           - show this help"