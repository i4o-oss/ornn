# Set default shell
set shell := ["bash", "-c"]

build:
	@echo "Building project..."
	go build -o bin/ornn

default:
	@just --list

install:
	@go install ./...

test:
	@go test -v ./...

vet:
	@go vet -v ./...
