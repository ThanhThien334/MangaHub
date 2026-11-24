.PHONY: build run-api run-catalog run-tui test fmt vet

build:
	go build -o bin/ ./cmd/...

run-api:
	go run ./cmd/api

run-catalog:
	go run ./cmd/catalog

run-tui:
	go run ./cmd/tui

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...