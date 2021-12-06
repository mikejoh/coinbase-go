CLIENT_VERSION?=$(shell git rev-parse --short HEAD)

client:
	mkdir -p ./bin
	go build -o ./bin/cb -ldflags "-X github.com/mikejoh/coinbase-go/coinbase.Version=$(CLIENT_VERSION)" ./cmd/cb

clean:
	go clean --cache
	rm -rf ./bin