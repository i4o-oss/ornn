# Set default shell
set shell := ["bash", "-c"]

build:
	@echo "Building project..."
	go build -o bin/ornn

default:
	@just --list

fmt:
	@go fmt ./...

install:
	@go install ./...

test:
	@go test -v ./...

tidy:
	@go mod tidy

vet:
	@go vet -v ./...
