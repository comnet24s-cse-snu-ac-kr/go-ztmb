BUILD=build

VERSION=0.1.0
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

all: build

build: build-linux-amd

build-darwin-arm:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD}/ztmb-doh-proxy cmd/ztmb-doh-proxy/*.go
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD}/ztmb-prover cmd/ztmb-prover/*.go
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD}/ztmb-wo-zkp cmd/ztmb-wo-zkp/*.go

build-linux-amd:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD}/ztmb-doh-proxy cmd/ztmb-doh-proxy/*.go
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD}/ztmb-prover cmd/ztmb-prover/*.go
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD}/ztmb-wo-zkp cmd/ztmb-wo-zkp/*.go

clean:
	rm -vf ${BUILD}/*

fmt:
	go fmt ./...

deps:
	go mod tidy

test:
	go test ${GO_FILES} -v

.PHONY: all build build-darwin-arm build-linux-amd clean fmt deps test
