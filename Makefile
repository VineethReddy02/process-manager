
build:
		go build -o bin/process-manager ./example

test:
		go test ./modules/*

govet:
		go vet ./modules/*

gofmt:
		gofmt -d .

golint:
		golangci-lint run

docker-build:
		docker build . -t vineeth97/process-manager:1.0

docker-push:
		docker push vineeth97/process-manager

all: build test govet gofmt golint docker-build docker-push



