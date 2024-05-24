BINARY := zkmbx
GO_FILES := $(wildcard src/*.go)
BUILD := build

VERSION=0.1.0
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

all: build

build:
	go build ${LDFLAGS} -o ${BINARY} ${GO_FILES}

build-all: build-arm build-amd

build-arm:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD}/${BINARY}-darwin-arm64 ${GO_FILES}

build-amd:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD}/${BINARY}-darwin-amd64 ${GO_FILES}

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)*

fmt:
	go fmt ./...

deps:
	go mod tidy

.PHONY: all build run clean test fmt lint deps help
