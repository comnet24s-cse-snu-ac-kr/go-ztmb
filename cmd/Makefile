PRV_BINARY := ztmb-prover
PRV_GO_FILES := $(wildcard ztmb-prover/*.go)

VRF_BINARY := ztmb-verifier
VRF_GO_FILES := $(wildcard ztmb-verifier/*.go)

VERSION=0.1.0
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

all: build

build:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${PRV_BINARY}-darwin-arm64 ${PRV_GO_FILES}
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${VRF_BINARY}-darwin-arm64 ${VRF_GO_FILES}

build-linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${PRV_BINARY}-linux-amd64 ${PRV_GO_FILES}
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${VRF_BINARY}-linux-amd64 ${VRF_GO_FILES}

clean:
	rm -f ${PRV_BINARY}* ${VRF_BINARY}*

fmt:
	go fmt ./...

deps:
	go mod tidy

test:
	go test ${PRV_GO_FILES} -v
	go test ${VRF_GO_FILES} -v

.PHONY: all build build-all build-arm build-amd install clean fmt deps test
