BINARY := ztmb-prover
GO_FILES := $(wildcard app/*.go)

VERSION=0.1.0
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

all: build

build: build-darwin-arm build-linux-amd

build-darwin-arm:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BINARY}-darwin-arm64 ${GO_FILES}

build-linux-amd:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-linux-amd64 ${GO_FILES}

clean:
	rm -f ${BINARY}*

fmt:
	go fmt ./...

deps:
	go mod tidy

test:
	go test ${GO_FILES} -v

.PHONY: all build build-darwin-arm build-linux-amd clean fmt deps test
