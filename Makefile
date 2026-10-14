.DEFAULT_GOAL := check

BINARY := homerow
PKG := ./...
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -X main.version=$(VERSION)

.PHONY: build
build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY) ./cmd/homerow

.PHONY: run
run: build
	./bin/$(BINARY)

.PHONY: fmt
fmt:
	@test -z "$$(gofmt -l .)" || { gofmt -l .; echo "run gofmt -w ."; exit 1; }

.PHONY: vet
vet:
	go vet $(PKG)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race $(PKG)

.PHONY: cover
cover:
	go test -coverprofile=coverage.out $(PKG)
	go tool cover -func=coverage.out | tail -1

.PHONY: check
check: fmt vet lint test

.PHONY: clean
clean:
	rm -rf bin coverage.out
