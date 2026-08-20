CMDPATH := ./cmd/cb
BINARY  := cb
BUILDDIR := bin

CLIENT_VERSION ?= $(shell git rev-parse --short HEAD)

LDFLAGS := -X github.com/mikejoh/coinbase-go.Version=$(CLIENT_VERSION)

.PHONY: test testcov dep vet lint clean build

## test: Run tests.
test:
	go test -v ./...

## testcov: Run tests with a coverage report.
testcov:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

## dep: Download and tidy module dependencies.
dep:
	go mod download
	go mod tidy

## vet: Run go vet.
vet:
	go vet ./...

## lint: Run golangci-lint.
lint:
	golangci-lint run -v --timeout=15m ./...

## clean: Remove build artifacts.
clean:
	go clean --cache
	rm -rf $(BUILDDIR)
	rm -f coverage.out

## build: Build the cb binary into $(BUILDDIR)/$(BINARY).
build:
	mkdir -p $(BUILDDIR)
	go build -o $(BUILDDIR)/$(BINARY) -ldflags "$(LDFLAGS)" $(CMDPATH)
