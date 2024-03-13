BINARY=mystocks
VERSION=0.1

run:
	go run main.go

build:
	go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$$(git rev-parse --short HEAD)" -o bin/$(BINARY)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$$(git rev-parse --short HEAD)" -o bin/$(BINARY)-linux-x86_64

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$$(git rev-parse --short HEAD)" -o bin/$(BINARY)-darwin-x86_64

build-mac-arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$$(git rev-parse --short HEAD)" -o bin/$(BINARY)-darwin-aarch64

changelog:
	echo "{}" > package.json && changelog generate && rm package.json

clean:
	rm -rf bin
