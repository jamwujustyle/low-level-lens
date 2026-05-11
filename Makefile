BACKEND_BIN := /bin/api
FRONTEND_DIR := ./interface

.DEFAULT_GOAL := check

.PHONY: build run test clean check watch dev

build:
	go build -o ${BACKEND_BIN} ./cmd/api/

run:
	go run ./cmd/api/

watch:
	air

test:
	go test ./...

clean:
	rm -rf bin/
	rm -rf $(FRONTEND_DIR)/.next
	rm -rf $(FRONTEND_DIR)/out

check:
	go fmt ./...
	go test ./...

dev:
	cd $(FRONTEND_DIR) && npm run dev

