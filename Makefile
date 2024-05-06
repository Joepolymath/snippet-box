.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o app ./cmd/web
.PHONY:build

run: build
	./app