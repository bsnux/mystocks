BINARY=mystocks
VERSION=0.1

run:
	go run main.go

build:
	go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$$(git rev-parse --short HEAD)" -o bin/$(BINARY)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$$(git rev-parse --short HEAD)" -o bin/$(BINARY)

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$$(git rev-parse --short HEAD)" -o bin/$(BINARY)

build-mac-arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$$(git rev-parse --short HEAD)" -o bin/$(BINARY)

clean:
	rm -rf bin
