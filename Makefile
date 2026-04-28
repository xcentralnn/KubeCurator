.PHONY: all fmt tidy test build verify clean

all: verify

fmt:
	go fmt ./...

tidy:
	go mod tidy

test:
	go test ./...

build:
	mkdir -p bin
	go build -o bin/curator ./cmd

verify: fmt tidy test build

clean:
	rm -rf bin
