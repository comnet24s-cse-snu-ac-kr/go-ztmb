# Makefile for a Go project

# Project variables
APP_NAME := logic
GO_FILES := $(wildcard *.go)

# Default target
all: build

# Build the Go project
build:
	@echo "Building the application..."
	go build -o $(APP_NAME) $(GO_FILES)

# Run the Go application
run: build
	@echo "Running the application..."
	./$(APP_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

# Test the Go project
test:
	@echo "Running tests..."
	go test ./...

# Format the Go source files
fmt:
	@echo "Formatting source files..."
	go fmt ./...

# Lint the Go source files
lint:
	@echo "Linting source files..."
	go vet ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Help message
help:
	@echo "Makefile for Go project"
	@echo
	@echo "Usage:"
	@echo "  make [target]"
	@echo
	@echo "Targets:"
	@echo "  all       - Build the application (default)"
	@echo "  build     - Build the application"
	@echo "  run       - Run the application"
	@echo "  clean     - Clean up build artifacts"
	@echo "  test      - Run tests"
	@echo "  fmt       - Format the Go source files"
	@echo "  lint      - Lint the Go source files"
	@echo "  deps      - Install dependencies"
	@echo "  help      - Show this help message"

.PHONY: all build run clean test fmt lint deps help
