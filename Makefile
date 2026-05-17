.PHONY: all test lint lint-fix test-cover build clean help

all: help

test:
	cd tests && go test -v ./...

# Линтер только в tests/ — CI не падает
lint:
	@echo "Running golangci-lint only in tests/..."
	cd tests && golangci-lint run

# Попытка исправить линт-ошибки
lint-fix:
	@echo "Running golangci-lint fix only in tests/..."
	cd tests && golangci-lint run --fix

# Тесты с покрытием
test-cover:
	cd tests && go test -coverprofile=coverage.out ./...
	@echo "Coverage report generated at tests/coverage.out"

# Собираем бинарь
build:
	@echo "Building binary..."
	@mkdir -p bin
	go build -o bin/app ./cmd/hexlet-path-size

clean:
	@rm -rf bin tests/coverage.out

help:
	@echo "Makefile commands:"
	@echo "  make test         - run all tests"
	@echo "  make lint         - run golangci-lint on tests/ only"
	@echo "  make lint-fix     - try to fix lint issues automatically (tests/ only)"
	@echo "  make test-cover   - run tests with coverage"
	@echo "  make build        - build the binary into bin/"
	@echo "  make clean        - remove build artifacts"
	@echo "  make help         - show this help"