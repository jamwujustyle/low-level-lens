.DEFAULT_GOAL := check

.PHONY: build run test clean check

build:
	go build -o bin/api ./cmd/api/

run:
	go run ./cmd/api/

test:
	go test ./...

clean:
	rm -rf bin/

check:
	go fmt ./...
	go test ./...
