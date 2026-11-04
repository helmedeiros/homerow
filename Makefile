.DEFAULT_GOAL := check

BINARY := homerow
PKG := ./...
COVER_MIN ?= 90
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
	go test -race -coverprofile=coverage.out $(PKG)

.PHONY: cover
cover: test
	@go tool cover -func=coverage.out | awk -v min=$(COVER_MIN) '/^total:/ { sub("%", "", $$3); print "coverage " $$3 "% (floor " min "%)"; if ($$3 + 0 < min) { print "coverage below floor"; exit 1 } }'

.PHONY: check
check: fmt vet lint cover

.PHONY: clean
clean:
	rm -rf bin coverage.out
